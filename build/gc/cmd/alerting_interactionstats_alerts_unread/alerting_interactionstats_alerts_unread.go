package alerting_interactionstats_alerts_unread

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
	Description = utils.FormatUsageDescription("alerting_interactionstats_alerts_unread", "SWAGGER_OVERRIDE_/api/v2/alerting/interactionstats/alerts/unread", )
	alerting_interactionstats_alerts_unreadCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("alerting_interactionstats_alerts_unread"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(alerting_interactionstats_alerts_unreadCmd)
}

func Cmdalerting_interactionstats_alerts_unread() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/alerting/interactionstats/alerts/unread", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Alerting", "/api/v2/alerting/interactionstats/alerts/unread")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/UnreadMetric&quot;
  }
}`)
	alerting_interactionstats_alerts_unreadCmd.AddCommand(getCmd)
	
	return alerting_interactionstats_alerts_unreadCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets user unread count of interaction stats alerts.",
	Long:  "Gets user unread count of interaction stats alerts.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/alerting/interactionstats/alerts/unread"

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
