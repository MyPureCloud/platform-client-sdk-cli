package telephony_providers_edges_trunks_metrics

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_trunks_metrics", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/trunks/{trunkId}/metrics", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/trunks/metrics", )
	telephony_providers_edges_trunks_metricsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_trunks_metrics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_trunks_metricsCmd)
}

func Cmdtelephony_providers_edges_trunks_metrics() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/trunks/{trunkId}/metrics", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/TrunkMetrics&quot;
  }
}`)
	telephony_providers_edges_trunks_metricsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "trunkIds", "", "Comma separated list of Trunk Id`s - REQUIRED")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/telephony/providers/edges/trunks/metrics", utils.FormatPermissions([]string{ "telephony:plugin:all",  })))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("trunkIds")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;type&quot; : &quot;array&quot;,
    &quot;items&quot; : {
      &quot;$ref&quot; : &quot;#/definitions/TrunkMetrics&quot;
    }
  }
}`)
	telephony_providers_edges_trunks_metricsCmd.AddCommand(listCmd)
	
	return telephony_providers_edges_trunks_metricsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [trunkId]",
	Short: "Get the trunk metrics.",
	Long:  "Get the trunk metrics.",
	Args:  utils.DetermineArgs([]string{ "trunkId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/trunks/{trunkId}/metrics"
		trunkId, args := args[0], args[1:]
		path = strings.Replace(path, "{trunkId}", fmt.Sprintf("%v", trunkId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get the metrics for a list of trunks.",
	Long:  "Get the metrics for a list of trunks.",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/trunks/metrics"

		trunkIds := utils.GetFlag(cmd.Flags(), "string", "trunkIds")
		if trunkIds != "" {
			queryParams["trunkIds"] = trunkIds
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
