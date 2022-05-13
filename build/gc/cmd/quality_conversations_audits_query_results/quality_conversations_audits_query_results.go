package quality_conversations_audits_query_results

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
	Description = utils.FormatUsageDescription("quality_conversations_audits_query_results", "SWAGGER_OVERRIDE_/api/v2/quality/conversations/audits/query/{transactionId}/results", )
	quality_conversations_audits_query_resultsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_conversations_audits_query_results"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_conversations_audits_query_resultsCmd)
}

func Cmdquality_conversations_audits_query_results() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "cursor", "", "Indicates where to resume query results (not required for first page)")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	utils.AddFlag(listCmd.Flags(), "[]string", "expand", "", "Which fields, if any, to expand Valid values: user")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/quality/conversations/audits/query/{transactionId}/results", utils.FormatPermissions([]string{ "audits:interactionDetails:view",  }), utils.GenerateDevCentreLink("GET", "Quality", "/api/v2/quality/conversations/audits/query/{transactionId}/results")))
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
	quality_conversations_audits_query_resultsCmd.AddCommand(listCmd)
	return quality_conversations_audits_query_resultsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [transactionId]",
	Short: "Get results of audit query",
	Long:  "Get results of audit query",
	Args:  utils.DetermineArgs([]string{ "transactionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/quality/conversations/audits/query/{transactionId}/results"
		transactionId, args := args[0], args[1:]
		path = strings.Replace(path, "{transactionId}", fmt.Sprintf("%v", transactionId), -1)

		cursor := utils.GetFlag(cmd.Flags(), "string", "cursor")
		if cursor != "" {
			queryParams["cursor"] = cursor
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
