package workforcemanagement_managementunits_activitycodes

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
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_activitycodes", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/activitycodes", )
	workforcemanagement_managementunits_activitycodesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_activitycodes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_activitycodesCmd)
}

func Cmdworkforcemanagement_managementunits_activitycodes() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/activitycodes", utils.FormatPermissions([]string{ "wfm:activityCode:add", "wfm:activityCode:delete", "wfm:activityCode:edit", "wfm:activityCode:view", "wfm:agent:edit", "wfm:agentSchedule:view", "wfm:agentTimeOffRequest:submit", "wfm:agent:view", "wfm:businessUnit:add", "wfm:businessUnit:delete", "wfm:businessUnit:edit", "wfm:businessUnit:view", "wfm:historicalAdherence:view", "wfm:intraday:view", "wfm:managementUnit:add", "wfm:managementUnit:delete", "wfm:managementUnit:edit", "wfm:managementUnit:view", "wfm:publishedSchedule:view", "wfm:realtimeAdherence:view", "wfm:schedule:add", "wfm:schedule:delete", "wfm:schedule:edit", "wfm:schedule:generate", "wfm:schedule:view", "wfm:shortTermForecast:add", "wfm:shortTermForecast:delete", "wfm:shortTermForecast:edit", "wfm:shortTermForecast:view", "wfm:timeOffRequest:add", "wfm:timeOffRequest:edit", "wfm:timeOffRequest:view", "wfm:workPlan:add", "wfm:workPlan:delete", "wfm:workPlan:edit", "wfm:workPlan:view", "wfm:workPlanRotation:add", "wfm:workPlanRotation:delete", "wfm:workPlanRotation:edit", "wfm:workPlanRotation:view",  })))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/ActivityCodeContainer&quot;
  }
}`)
	workforcemanagement_managementunits_activitycodesCmd.AddCommand(getCmd)
	
	return workforcemanagement_managementunits_activitycodesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [managementUnitId]",
	Short: "Get activity codes",
	Long:  "Get activity codes",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/activitycodes"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)

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
