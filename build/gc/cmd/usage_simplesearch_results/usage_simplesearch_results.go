package usage_simplesearch_results

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
	Description = utils.FormatUsageDescription("usage_simplesearch_results", "SWAGGER_OVERRIDE_/api/v2/usage/simplesearch/{executionId}/results", )
	usage_simplesearch_resultsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage_simplesearch_results"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usage_simplesearch_resultsCmd)
}

func Cmdusage_simplesearch_results() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "after", "", "The cursor that points to the end of the set of entities that has been returned")
	utils.AddFlag(getCmd.Flags(), "int", "pageSize", "", "The max number of entities to be returned per request. Maximum page size of 1000")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/usage/simplesearch/{executionId}/results", utils.FormatPermissions([]string{ "oauth:client:view", "usage:simpleSearch:view",  }), utils.GenerateDevCentreLink("GET", "Usage", "/api/v2/usage/simplesearch/{executionId}/results")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/ApiUsageQueryResult"
      }
    }
  }
}`)
	usage_simplesearch_resultsCmd.AddCommand(getCmd)
	return usage_simplesearch_resultsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var getCmd = &cobra.Command{
	Use:   "get [executionId]",
	Short: "Get the results of a usage search. Number of records to be returned is limited to 20,000 results.",
	Long:  "Get the results of a usage search. Number of records to be returned is limited to 20,000 results.",
	Args:  utils.DetermineArgs([]string{ "executionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/usage/simplesearch/{executionId}/results"
		executionId, args := args[0], args[1:]
		path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionId), -1)

		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
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
