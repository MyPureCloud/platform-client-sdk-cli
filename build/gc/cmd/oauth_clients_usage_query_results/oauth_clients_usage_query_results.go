package oauth_clients_usage_query_results

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
	Description = utils.FormatUsageDescription("oauth_clients_usage_query_results", "SWAGGER_OVERRIDE_/api/v2/oauth/clients/{clientId}/usage/query/results", )
	oauth_clients_usage_query_resultsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("oauth_clients_usage_query_results"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(oauth_clients_usage_query_resultsCmd)
}

func Cmdoauth_clients_usage_query_results() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}", utils.FormatPermissions([]string{ "oauth:client:view",  }), utils.GenerateDevCentreLink("GET", "OAuth", "/api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}")))
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
	oauth_clients_usage_query_resultsCmd.AddCommand(getCmd)
	return oauth_clients_usage_query_resultsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [executionId] [clientId]",
	Short: "Get the results of a usage query",
	Long:  "Get the results of a usage query",
	Args:  utils.DetermineArgs([]string{ "executionId", "clientId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/oauth/clients/{clientId}/usage/query/results/{executionId}"
		executionId, args := args[0], args[1:]
		path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionId), -1)
		clientId, args := args[0], args[1:]
		path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
