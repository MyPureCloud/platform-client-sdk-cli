package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/restclient"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
)

//CommandService holds the method signatures for all common Command object invocations
type CommandService interface {
	Get(uri string) (string, error)
	List(uri string) (string, error)
	Stream(uri string) (string, error)
	Post(uri string, payload string) (string, error)
	Patch(uri string, payload string) (string, error)
	Put(uri string, payload string) (string, error)
	Delete(uri string) (string, error)
	DetermineAction(httpMethod string, uri string, flags *pflag.FlagSet) func(retryConfiguration *retry.RetryConfiguration) (string, error)
}

type commandService struct {
	cmd       *cobra.Command
	startTime time.Time
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
	c.traceStart(http.MethodGet, uri, "")
	response, err := restClient.Get(uri)
	if err == nil {
		c.traceEnd()
		return response, nil
	}

	err = reAuthenticateIfNecessary(config, err)
	if err != nil {
		return "", err
	}

	return c.Get(uri)
}

func (c *commandService) Stream(uri string) (string, error) {
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
		return c.Stream(uri)
	}

	utils.Render(data)

	firstPage := &models.Entities{}
	json.Unmarshal([]byte(data), firstPage)
	if firstPage.PageCount == 1 || firstPage.PageNumber >= firstPage.PageCount {
		return "", nil
	}

	pagedURI := uri
	for x := firstPage.PageNumber + 1; x <= firstPage.PageCount; x++ {
		pagedURI = updatePageNumber(pagedURI, x)
		logger.Info("Paginating with URI: ", pagedURI)
		retryFunc := retry.Retry(pagedURI, restClient.Get)
		data, err = retryFunc(&retry.RetryConfiguration{
			RetryWaitMax: 1000 * time.Second,
			RetryWaitMin: 1000 * time.Second,
			RetryMax:     100,
		})
		if err != nil {
			return "", err
		}

		utils.Render(data)
	}

	return "", nil
}

func (c *commandService) List(uri string) (string, error) {
	profileName, _ := c.cmd.Root().Flags().GetString("profile")
	config, err := configGetConfig(profileName)
	if err != nil {
		return "", err
	}

	restClient := restclientNewRESTClient(config)

	//Looks up first page
	c.traceStart(http.MethodGet, uri, "")
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
		for x := firstPage.PageNumber + 1; x <= firstPage.PageCount; x++ {
			pagedURI = updatePageNumber(pagedURI, x)
			logger.Info("Paginating with URI: ", pagedURI)
			c.traceProgress(pagedURI)
			retryFunc := retry.Retry(pagedURI, restClient.Get)
			data, err = retryFunc(&retry.RetryConfiguration{
				RetryWaitMax: 1000 * time.Second,
				RetryWaitMin: 1000 * time.Second,
				RetryMax:     100,
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

	c.traceEnd()

	return finalJSONString, nil
}

func updatePageNumber(pagedURI string, index int) string {
	if strings.Contains(pagedURI, "pageNumber=") {
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

	c.traceStart(method, uri, payload)
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
		c.traceEnd()
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

func (c *commandService) DetermineAction(httpMethod string, uri string, flags *pflag.FlagSet) func(retryConfiguration *retry.RetryConfiguration) (string, error) {
	logger.InitLogger(c.cmd)
	switch httpMethod {
	case http.MethodGet:
		if flags == nil {
			return retry.Retry(uri, c.Get)
		}
		// These flags will be false if they're not available on the command (simple GETs) or if they haven't been set on a paginatable command
		autoPaginate, _ := flags.GetBool("autopaginate")
		stream, _ := flags.GetBool("stream")
		if !autoPaginate && !stream {
			return retry.Retry(uri, c.Get)
		}
		// Stream if the user just sets stream or stream and autopaginate
		if stream {
			return retry.Retry(uri, c.Stream)
		}
		return retry.Retry(uri, c.List)
	case http.MethodPatch:
		if flags.Lookup("file") == nil {
			return retry.RetryWithData(uri, "", c.Patch)
		}
		return retry.RetryWithData(uri, utils.ResolveInputData(c.cmd), c.Patch)
	case http.MethodPost:
		if flags.Lookup("file") == nil {
			return retry.RetryWithData(uri, "", c.Post)
		}
		return retry.RetryWithData(uri, utils.ResolveInputData(c.cmd), c.Post)
	case http.MethodPut:
		if flags.Lookup("file") == nil {
			return retry.RetryWithData(uri, "", c.Put)
		}
		return retry.RetryWithData(uri, utils.ResolveInputData(c.cmd), c.Put)
	case http.MethodDelete:
		return retry.Retry(uri, c.Delete)
	}
	return nil
}

func (c *commandService) traceStart(method, uri, data string) {
	traceProgress, _ := c.cmd.Root().Flags().GetBool("indicateprogress")
	if traceProgress {
		c.startTime = time.Now()
		if data != "" {
			logger.Tracef("Command started at: %v. Method: %v, Path: %v, Data: %v\n", c.startTime.Format(time.RFC3339Nano), method, uri, data)
		} else {
			logger.Tracef("Command started at: %v. Method: %v, Path: %v\n", c.startTime.Format(time.RFC3339Nano), method, uri)
		}
	}
}

func (c *commandService) traceProgress(pagedURI string) {
	traceProgress, _ := c.cmd.Root().Flags().GetBool("indicateprogress")
	if traceProgress {
		logger.Tracef("Paginating with path: %v\n", pagedURI)
	}
}

func (c *commandService) traceEnd() {
	traceProgress, _ := c.cmd.Root().Flags().GetBool("indicateprogress")
	if traceProgress {
		endTime := time.Now()
		logger.Tracef("Command finished at: %v. Time taken: %v\n", endTime.Format(time.RFC3339Nano), endTime.Sub(c.startTime))
	}
}
