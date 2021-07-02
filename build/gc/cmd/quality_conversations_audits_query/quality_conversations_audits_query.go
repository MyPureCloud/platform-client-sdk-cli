package quality_conversations_audits_query

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
	Description = utils.FormatUsageDescription("quality_conversations_audits_query", "SWAGGER_OVERRIDE_/api/v2/quality/conversations/audits/query", "SWAGGER_OVERRIDE_/api/v2/quality/conversations/audits/query", )
	quality_conversations_audits_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_conversations_audits_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_conversations_audits_queryCmd)
}

func Cmdquality_conversations_audits_query() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/quality/conversations/audits/query", utils.FormatPermissions([]string{ "audits:interactionDetails:view",  }), utils.GenerateDevCentreLink("POST", "Quality", "/api/v2/quality/conversations/audits/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;query&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/QMAuditQueryRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/QualityAuditQueryExecutionStatusResponse&quot;
  }
}`)
	quality_conversations_audits_queryCmd.AddCommand(createCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/quality/conversations/audits/query/{transactionId}", utils.FormatPermissions([]string{ "audits:interactionDetails:view",  }), utils.GenerateDevCentreLink("GET", "Quality", "/api/v2/quality/conversations/audits/query/{transactionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;Query execution completed.&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/QualityAuditQueryExecutionStatusResponse&quot;
  }
}`)
	quality_conversations_audits_queryCmd.AddCommand(getCmd)
	
	return quality_conversations_audits_queryCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create audit query execution",
	Long:  "Create audit query execution",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/quality/conversations/audits/query"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
var getCmd = &cobra.Command{
	Use:   "get [transactionId]",
	Short: "Get status of audit query execution",
	Long:  "Get status of audit query execution",
	Args:  utils.DetermineArgs([]string{ "transactionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/quality/conversations/audits/query/{transactionId}"
		transactionId, args := args[0], args[1:]
		path = strings.Replace(path, "{transactionId}", fmt.Sprintf("%v", transactionId), -1)

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
