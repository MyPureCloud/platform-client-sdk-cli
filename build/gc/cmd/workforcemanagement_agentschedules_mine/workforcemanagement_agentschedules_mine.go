package workforcemanagement_agentschedules_mine

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
	Description = utils.FormatUsageDescription("workforcemanagement_agentschedules_mine", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agentschedules/mine", )
	workforcemanagement_agentschedules_mineCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agentschedules_mine"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agentschedules_mineCmd)
}

func Cmdworkforcemanagement_agentschedules_mine() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/agentschedules/mine", utils.FormatPermissions([]string{ "wfm:agentSchedule:view",  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;body&quot;,
  &quot;required&quot; : false,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/BuGetCurrentAgentScheduleRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/BuCurrentAgentScheduleSearchResponse&quot;
  }
}`)
	workforcemanagement_agentschedules_mineCmd.AddCommand(createCmd)
	
	return workforcemanagement_agentschedules_mineCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Get published schedule for the current user",
	Long:  "Get published schedule for the current user",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/agentschedules/mine"

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
