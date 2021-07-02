package telephony_providers_edges_sites_rebalance

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_sites_rebalance", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/sites/{siteId}/rebalance", )
	telephony_providers_edges_sites_rebalanceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_sites_rebalance"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_sites_rebalanceCmd)
}

func Cmdtelephony_providers_edges_sites_rebalance() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/telephony/providers/edges/sites/{siteId}/rebalance", utils.FormatPermissions([]string{ "telephony:plugin:all",  }), utils.GenerateDevCentreLink("POST", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/sites/{siteId}/rebalance")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Accepted - Processing the Rebalance&quot;
}`)
	telephony_providers_edges_sites_rebalanceCmd.AddCommand(createCmd)
	
	return telephony_providers_edges_sites_rebalanceCmd
}

var createCmd = &cobra.Command{
	Use:   "create [siteId]",
	Short: "Triggers the rebalance operation.",
	Long:  "Triggers the rebalance operation.",
	Args:  utils.DetermineArgs([]string{ "siteId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/sites/{siteId}/rebalance"
		siteId, args := args[0], args[1:]
		path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)

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
