package infrastructureascode_accelerators

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("infrastructureascode_accelerators", "SWAGGER_OVERRIDE_/api/v2/infrastructureascode/accelerators", "SWAGGER_OVERRIDE_/api/v2/infrastructureascode/accelerators", )
	infrastructureascode_acceleratorsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("infrastructureascode_accelerators"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(infrastructureascode_acceleratorsCmd)
}

func Cmdinfrastructureascode_accelerators() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "preferredLanguage", "en-US", "Preferred Language Valid values: ar, cs, da, de, en-US, es, fi, fr, it, iw, ko, ja, nl, no, pl, pt-BR, pt-PT, sv, th, tr, zh-CN, zh-TW")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/infrastructureascode/accelerators/{acceleratorId}", utils.FormatPermissions([]string{ "infrastructureascode:accelerator:view",  }), utils.GenerateDevCentreLink("GET", "Infrastructure as Code", "/api/v2/infrastructureascode/accelerators/{acceleratorId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AcceleratorSpecification"
      }
    }
  }
}`)
	infrastructureascode_acceleratorsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "The total page size requested")
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "The page number requested")
	utils.AddFlag(listCmd.Flags(), "string", "sortBy", "", "variable name requested to sort by")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "asc", "Sort order Valid values: asc, desc")
	utils.AddFlag(listCmd.Flags(), "string", "name", "", "Filter by name")
	utils.AddFlag(listCmd.Flags(), "string", "description", "", "Filter by description")
	utils.AddFlag(listCmd.Flags(), "string", "origin", "", "Filter by origin Valid values: community, partner, genesys")
	utils.AddFlag(listCmd.Flags(), "string", "varType", "", "Filter by type Valid values: module, accelerator, blueprint")
	utils.AddFlag(listCmd.Flags(), "string", "classification", "", "Filter by classification")
	utils.AddFlag(listCmd.Flags(), "string", "tags", "", "Filter by tags")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/infrastructureascode/accelerators", utils.FormatPermissions([]string{ "infrastructureascode:accelerator:view",  }), utils.GenerateDevCentreLink("GET", "Infrastructure as Code", "/api/v2/infrastructureascode/accelerators")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	infrastructureascode_acceleratorsCmd.AddCommand(listCmd)
	return infrastructureascode_acceleratorsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [acceleratorId]",
	Short: "Get information about an accelerator",
	Long:  "Get information about an accelerator",
	Args:  utils.DetermineArgs([]string{ "acceleratorId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/infrastructureascode/accelerators/{acceleratorId}"
		acceleratorId, args := args[0], args[1:]
		path = strings.Replace(path, "{acceleratorId}", fmt.Sprintf("%v", acceleratorId), -1)

		preferredLanguage := utils.GetFlag(cmd.Flags(), "string", "preferredLanguage")
		if preferredLanguage != "" {
			queryParams["preferredLanguage"] = preferredLanguage
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of available accelerators",
	Long:  "Get a list of available accelerators",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/infrastructureascode/accelerators"

		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		sortBy := utils.GetFlag(cmd.Flags(), "string", "sortBy")
		if sortBy != "" {
			queryParams["sortBy"] = sortBy
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		name := utils.GetFlag(cmd.Flags(), "string", "name")
		if name != "" {
			queryParams["name"] = name
		}
		description := utils.GetFlag(cmd.Flags(), "string", "description")
		if description != "" {
			queryParams["description"] = description
		}
		origin := utils.GetFlag(cmd.Flags(), "string", "origin")
		if origin != "" {
			queryParams["origin"] = origin
		}
		varType := utils.GetFlag(cmd.Flags(), "string", "varType")
		if varType != "" {
			queryParams["type"] = varType
		}
		classification := utils.GetFlag(cmd.Flags(), "string", "classification")
		if classification != "" {
			queryParams["classification"] = classification
		}
		tags := utils.GetFlag(cmd.Flags(), "string", "tags")
		if tags != "" {
			queryParams["tags"] = tags
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "list"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
