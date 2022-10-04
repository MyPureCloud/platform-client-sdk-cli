package oauth_clients_usage_summary

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
	Description = utils.FormatUsageDescription("oauth_clients_usage_summary", "SWAGGER_OVERRIDE_/api/v2/oauth/clients/{clientId}/usage/summary", )
	oauth_clients_usage_summaryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("oauth_clients_usage_summary"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(oauth_clients_usage_summaryCmd)
}

func Cmdoauth_clients_usage_summary() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "days", "7", "Previous number of days to query")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/oauth/clients/{clientId}/usage/summary", utils.FormatPermissions([]string{ "oauth:client:view",  }), utils.GenerateDevCentreLink("GET", "OAuth", "/api/v2/oauth/clients/{clientId}/usage/summary")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UsageExecutionResult"
      }
    }
  }
}`)
	oauth_clients_usage_summaryCmd.AddCommand(getCmd)
	return oauth_clients_usage_summaryCmd
}

var getCmd = &cobra.Command{
	Use:   "get [clientId]",
	Short: "Get a summary of OAuth client API usage",
	Long:  "Get a summary of OAuth client API usage",
	Args:  utils.DetermineArgs([]string{ "clientId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/oauth/clients/{clientId}/usage/summary"
		clientId, args := args[0], args[1:]
		path = strings.Replace(path, "{clientId}", fmt.Sprintf("%v", clientId), -1)

		days := utils.GetFlag(cmd.Flags(), "string", "days")
		if days != "" {
			queryParams["days"] = days
		}
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
