package workforcemanagement_businessunits_weeks_schedules_headcountforecast

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_schedules_headcountforecast", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/headcountforecast", )
	workforcemanagement_businessunits_weeks_schedules_headcountforecastCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_schedules_headcountforecast"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_weeks_schedules_headcountforecastCmd)
}

func Cmdworkforcemanagement_businessunits_weeks_schedules_headcountforecast() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "bool", "forceDownload", "", "Whether to force the result to come via download url.  For testing purposes only")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/headcountforecast", utils.FormatPermissions([]string{ "wfm:schedule:view", "wfm:publishedSchedule:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/headcountforecast")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/BuHeadcountForecastResponse&quot;
  }
}`)
	workforcemanagement_businessunits_weeks_schedules_headcountforecastCmd.AddCommand(getCmd)
	
	return workforcemanagement_businessunits_weeks_schedules_headcountforecastCmd
}

var getCmd = &cobra.Command{
	Use:   "get [businessUnitId] [weekId] [scheduleId]",
	Short: "Get the headcount forecast by planning group for the schedule",
	Long:  "Get the headcount forecast by planning group for the schedule",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "weekId", "scheduleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/headcountforecast"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
		weekId, args := args[0], args[1:]
		path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
		scheduleId, args := args[0], args[1:]
		path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)

		forceDownload := utils.GetFlag(cmd.Flags(), "bool", "forceDownload")
		if forceDownload != "" {
			queryParams["forceDownload"] = forceDownload
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
