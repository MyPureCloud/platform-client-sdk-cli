package workforcemanagement_managementunits_weeks_schedules

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
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_weeks_schedules", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules", )
	workforcemanagement_managementunits_weeks_schedulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_weeks_schedules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_weeks_schedulesCmd)
}

func Cmdworkforcemanagement_managementunits_weeks_schedules() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "expand", "", "Which fields, if any, to expand Valid values: generationResults, headcountForecast")
	utils.AddFlag(getCmd.Flags(), "bool", "forceDownloadService", "", "Force the result of this operation to be sent via download service.  For testing/app development purposes")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}", utils.FormatPermissions([]string{ "wfm:publishedSchedule:view", "wfm:schedule:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/WeekScheduleResponse"
  }
}`)
	workforcemanagement_managementunits_weeks_schedulesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "bool", "includeOnlyPublished", "", "Return only published schedules")
	utils.AddFlag(listCmd.Flags(), "string", "earliestWeekDate", "", "The start date of the earliest week to query in yyyy-MM-dd format")
	utils.AddFlag(listCmd.Flags(), "string", "latestWeekDate", "", "The start date of the latest week to query in yyyy-MM-dd format")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules", utils.FormatPermissions([]string{ "wfm:publishedSchedule:view", "wfm:schedule:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/WeekScheduleListResponse"
  }
}`)
	workforcemanagement_managementunits_weeks_schedulesCmd.AddCommand(listCmd)
	
	return workforcemanagement_managementunits_weeks_schedulesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [managementUnitId] [weekId] [scheduleId]",
	Short: "Deprecated.  Use the equivalent business unit resource instead. Get a week schedule",
	Long:  "Deprecated.  Use the equivalent business unit resource instead. Get a week schedule",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "weekId", "scheduleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		weekId, args := args[0], args[1:]
		path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
		scheduleId, args := args[0], args[1:]
		path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)


		expand := utils.GetFlag(cmd.Flags(), "string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		forceDownloadService := utils.GetFlag(cmd.Flags(), "bool", "forceDownloadService")
		if forceDownloadService != "" {
			queryParams["forceDownloadService"] = forceDownloadService
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
var listCmd = &cobra.Command{
	Use:   "list [managementUnitId] [weekId]",
	Short: "Deprecated.  Use the equivalent business unit resource instead. Get the list of schedules in a week in management unit",
	Long:  "Deprecated.  Use the equivalent business unit resource instead. Get the list of schedules in a week in management unit",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "weekId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		weekId, args := args[0], args[1:]
		path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)


		includeOnlyPublished := utils.GetFlag(cmd.Flags(), "bool", "includeOnlyPublished")
		if includeOnlyPublished != "" {
			queryParams["includeOnlyPublished"] = includeOnlyPublished
		}
		earliestWeekDate := utils.GetFlag(cmd.Flags(), "string", "earliestWeekDate")
		if earliestWeekDate != "" {
			queryParams["earliestWeekDate"] = earliestWeekDate
		}
		latestWeekDate := utils.GetFlag(cmd.Flags(), "string", "latestWeekDate")
		if latestWeekDate != "" {
			queryParams["latestWeekDate"] = latestWeekDate
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
