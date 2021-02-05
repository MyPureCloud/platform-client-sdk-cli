package services

import (
	"encoding/json"
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/restclient"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

//CommandService holds the method signatures for all common Command object invocations
type CommandService interface {
	Get(uri string) (string, error)
	List(uri string) (string, error)
	Post(uri string, payload string) (string, error)
	Patch(uri string, payload string) (string, error)
	Put(uri string, payload string) (string, error)
	Delete(uri string) (string, error)
	DetermineAction(httpMethod string, operationId string, uri string, originalURI string) func(retryConfiguration *retry.RetryConfiguration) (string, error)
}

type commandService struct {
	cmd *cobra.Command
}

// The following functions are added as variables to allow reassignment to mock functions in unit tests
var (
	configGetConfig         = config.GetConfig
	restclientNewRESTClient = restclient.NewRESTClient
	// This will be set if the application silently re-authenticates for a 401 error
	hasReAuthenticated bool
)

//NewCommandService initializes a new command Service object
func NewCommandService(cmd *cobra.Command) *commandService {
	return &commandService{cmd: cmd}
}

func (c *commandService) Get(uri string) (string, error) {
	profileName, _ := c.cmd.Root().Flags().GetString("profile")
	config, err := configGetConfig(profileName)
	if err != nil {
		return "", err
	}

	restClient := restclientNewRESTClient(config)
	response, err := restClient.Get(uri)
	if err == nil {
		return response, nil
	}

	err = reAuthenticateIfNecessary(config, err)
	if err != nil {
		return "", err
	}

	return c.Get(uri)
}

func (c *commandService) List(uri string) (string, error) {
	profileName, _ := c.cmd.Root().Flags().GetString("profile")
	config, err := configGetConfig(profileName)
	if err != nil {
		return "", err
	}

	restClient := restclientNewRESTClient(config)

	//Looks up first page
	data, err := restClient.Get(uri)
	if err != nil {
		err = reAuthenticateIfNecessary(config, err)
		if err != nil {
			return "", err
		}
		return c.List(uri)
	}

	firstPage := &models.Entities{}
	json.Unmarshal([]byte(data), firstPage)

	//Allocate the total results based on the page count
	totalResults := make([]string, 0)

	//Appends the individual records from Entities into the array
	for _, val := range firstPage.Entities {
		totalResults = append(totalResults, string(val))
	}

    //Looks up the rest of the pages
	if firstPage.PageCount > 1 {
		pagedURI := uri
		for x := 2; x <= firstPage.PageCount; x++ {
			pagedURI = updatePageNumber(pagedURI, x)
			logger.Info("Paginating with URI: ", pagedURI)
			retryFunc := retry.Retry(pagedURI, restClient.Get)
			data, err = retryFunc(&retry.RetryConfiguration{
				MaxRetryTimeSec: 20,
				MaxRetriesBeforeQuitting: 10,
			})
			if err != nil {
				return "", err
			}

			pageData := &models.Entities{}
			json.Unmarshal([]byte(data), pageData)

			for _, val := range pageData.Entities {
				totalResults = append(totalResults, string(val))
			}
		}
	}

	//Convert the data into one big string
	var finalJSONString string
	switch len(totalResults) {
	case 0:
		finalJSONString = "[]"
	case 1:
		finalJSONString = fmt.Sprintf("[%s]", totalResults[0])
	default:
		finalJSONString = fmt.Sprintf("[%s]", strings.Join(totalResults, ","))
	}

	return finalJSONString, nil
}

func updatePageNumber(pagedURI string, index int) string {
	if strings.Contains(pagedURI,"pageNumber=") {
		re := regexp.MustCompile("pageNumber=([0-9]+)")
		result := re.FindStringSubmatch(pagedURI)
		pageNumber, _ := strconv.Atoi(result[1])
		pageNumber++
		pagedURI = strings.Replace(pagedURI, result[0], fmt.Sprintf("pageNumber=%v", pageNumber), 1)
	} else {
		if strings.Contains(pagedURI, "?") {
			pagedURI = fmt.Sprintf("%s&pageNumber=%d", pagedURI, index)
		} else {
			pagedURI = fmt.Sprintf("%s?pageNumber=%d", pagedURI, index)
		}
	}

	return pagedURI
}

func (c *commandService) Post(uri string, payload string) (string, error) {
	return c.upsert(http.MethodPost, uri, payload)
}

func (c *commandService) Patch(uri string, payload string) (string, error) {
	return c.upsert(http.MethodPatch, uri, payload)
}

func (c *commandService) Put(uri string, payload string) (string, error) {
	return c.upsert(http.MethodPut, uri, payload)
}

func (c *commandService) Delete(uri string) (string, error) {
	return c.upsert(http.MethodDelete, uri, "")
}

func (c *commandService) upsert(method string, uri string, payload string) (string, error) {
	profileName, _ := c.cmd.Root().Flags().GetString("profile")
	config, err := configGetConfig(profileName)
	if err != nil {
		return "", err
	}

	restClient := restclientNewRESTClient(config)

	var response string

	switch method {
	case http.MethodPost:
		response, err = restClient.Post(uri, payload)
	case http.MethodPut:
		response, err = restClient.Put(uri, payload)
	case http.MethodPatch:
		response, err = restClient.Patch(uri, payload)
	case http.MethodDelete:
		response, err = restClient.Delete(uri)
	default:
		logger.Fatalf("Unable to resolve the http verb: %v", method)
	}

	if err == nil {
		return response, nil
	}

	err = reAuthenticateIfNecessary(config, err)
	if err != nil {
		return "", err
	}

	return c.upsert(method, uri, payload)
}

func reAuthenticateIfNecessary(config config.Configuration, err error) error {
	if hasReAuthenticated {
		logger.Warn("Have already re-authenticated. Will not authenticate again")
		return err
	}

	if e, ok := err.(models.HttpStatusError); ok && e.StatusCode == http.StatusUnauthorized {
		logger.Info("Received HTTP 401 error, re-authenticating")
		_, err = restclient.ReAuthenticate(config)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	hasReAuthenticated = true

	return nil
}

func (c *commandService) DetermineAction(httpMethod string, operationId string, uri string, originalURI string) func(retryConfiguration *retry.RetryConfiguration) (string, error) {
	switch httpMethod {
	case http.MethodGet:
		listOverrides := make(map[string]int)
		// Add overrides here for resources with custom operationIds requiring pagination
		listOverrides["/api/v2/routing/queues/{queueId}/users"] = 1
		listOverrides["/api/v2/users/{userId}/queues"] = 1
		listOverrides["/api/v2/groups/{groupId}/members"] = 1

		_, ok := listOverrides[originalURI]
		if operationId == "list" || ok {
			return retry.Retry(uri, c.List)
		} else {
			return retry.Retry(uri, c.Get)
		} 
	case http.MethodPatch:
		return retry.RetryWithData(uri, utils.ResolveInputData(c.cmd), c.Patch)
	case http.MethodPost:
		return retry.RetryWithData(uri, utils.ResolveInputData(c.cmd), c.Post)
	case http.MethodPut:
		return retry.RetryWithData(uri, utils.ResolveInputData(c.cmd), c.Put)
	case http.MethodDelete:
		return retry.Retry(uri, c.Delete)
	}
	return nil
}
