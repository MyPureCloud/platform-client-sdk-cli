package integrations_botconnector_bots_summaries

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
	Description = utils.FormatUsageDescription("integrations_botconnector_bots_summaries", "SWAGGER_OVERRIDE_/api/v2/integrations/botconnector/{integrationId}/bots/summaries", )
	integrations_botconnector_bots_summariesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_botconnector_bots_summaries"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_botconnector_bots_summariesCmd)
}

func Cmdintegrations_botconnector_bots_summaries() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "Page number")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "Page size")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/integrations/botconnector/{integrationId}/bots/summaries", utils.FormatPermissions([]string{ "integration:botconnector:view",  }), utils.GenerateDevCentreLink("GET", "Integrations", "/api/v2/integrations/botconnector/{integrationId}/bots/summaries")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/SWAGGER_OVERRIDE_list&quot;
  }
}`)
	integrations_botconnector_bots_summariesCmd.AddCommand(listCmd)
	
	return integrations_botconnector_bots_summariesCmd
}

var listCmd = &cobra.Command{
	Use:   "list [integrationId]",
	Short: "Get a summary list of botConnector bots for this integration",
	Long:  "Get a summary list of botConnector bots for this integration",
	Args:  utils.DetermineArgs([]string{ "integrationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/integrations/botconnector/{integrationId}/bots/summaries"
		integrationId, args := args[0], args[1:]
		path = strings.Replace(path, "{integrationId}", fmt.Sprintf("%v", integrationId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
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
