package telephony_providers_edges_autoscalinggroups_capacity

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
	Description = utils.FormatUsageDescription("telephony_providers_edges_autoscalinggroups_capacity", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/autoscalinggroups/{asgId}/capacity", )
	telephony_providers_edges_autoscalinggroups_capacityCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_autoscalinggroups_capacity"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_autoscalinggroups_capacityCmd)
}

func Cmdtelephony_providers_edges_autoscalinggroups_capacity() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/telephony/providers/edges/autoscalinggroups/{asgId}/capacity", utils.FormatPermissions([]string{ "telephony:plugin:all", "internal:edge:edit",  }), utils.GenerateDevCentreLink("PATCH", "Telephony Providers Edge", "/api/v2/telephony/providers/edges/autoscalinggroups/{asgId}/capacity")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;AsgScaleRequest&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AsgScaleRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ScaleASGResponse&quot;
  }
}`)
	telephony_providers_edges_autoscalinggroups_capacityCmd.AddCommand(updateCmd)
	
	return telephony_providers_edges_autoscalinggroups_capacityCmd
}

var updateCmd = &cobra.Command{
	Use:   "update [asgId]",
	Short: "Scales the ASG to match the desired capacity",
	Long:  "Scales the ASG to match the desired capacity",
	Args:  utils.DetermineArgs([]string{ "asgId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/telephony/providers/edges/autoscalinggroups/{asgId}/capacity"
		asgId, args := args[0], args[1:]
		path = strings.Replace(path, "{asgId}", fmt.Sprintf("%v", asgId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
