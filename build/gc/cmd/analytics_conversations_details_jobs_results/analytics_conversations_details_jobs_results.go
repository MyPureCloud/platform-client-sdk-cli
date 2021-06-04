package analytics_conversations_details_jobs_results

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("analytics_conversations_details_jobs_results", "SWAGGER_OVERRIDE_/api/v2/analytics/conversations/details/jobs/{jobId}/results", )
	analytics_conversations_details_jobs_resultsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_conversations_details_jobs_results"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_conversations_details_jobs_resultsCmd)
}

func Cmdanalytics_conversations_details_jobs_results() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "cursor", "", "Indicates where to resume query results (not required for first page)")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "", "The desired maximum number of results")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/analytics/conversations/details/jobs/{jobId}/results", utils.FormatPermissions([]string{ "analytics:conversationDetail:view",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	analytics_conversations_details_jobs_resultsCmd.AddCommand(listCmd)
	
	return analytics_conversations_details_jobs_resultsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [jobId]",
	Short: "Fetch a page of results for an async query",
	Long:  "Fetch a page of results for an async query",
	Args:  utils.DetermineArgs([]string{ "jobId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/analytics/conversations/details/jobs/{jobId}/results"
		jobId, args := args[0], args[1:]
		path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)

		cursor := utils.GetFlag(cmd.Flags(), "string", "cursor")
		if cursor != "" {
			queryParams["cursor"] = cursor
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
