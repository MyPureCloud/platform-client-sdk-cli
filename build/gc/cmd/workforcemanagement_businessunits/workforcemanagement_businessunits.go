package workforcemanagement_businessunits

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits", )
	workforcemanagement_businessunitsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunitsCmd)
}

func Cmdworkforcemanagement_businessunits() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "includeSchedulingDefaultMessageSeverities", "", "Whether to include scheduling default message severities")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/businessunits", utils.FormatPermissions([]string{ "wfm:businessUnit:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/businessunits")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CreateBusinessUnitRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BusinessUnitResponse"
      }
    }
  }
}`)
	workforcemanagement_businessunitsCmd.AddCommand(createCmd)

	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/workforcemanagement/businessunits/{businessUnitId}", utils.FormatPermissions([]string{ "wfm:businessUnit:delete",  }), utils.GenerateDevCentreLink("DELETE", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  "description" : "The business unit was successfully deleted",
  "content" : { }
}`)
	workforcemanagement_businessunitsCmd.AddCommand(deleteCmd)

	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "Include to access additional data on the business unit Valid values: settings, settings.timeZone, settings.startDayOfWeek, settings.shortTermForecasting, settings.scheduling, settings.notifications.scheduling, settings.learning, settings.coaching")
	utils.AddFlag(getCmd.Flags(), "bool", "includeSchedulingDefaultMessageSeverities", "", "Whether to include scheduling default message severities")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits/{businessUnitId}", utils.FormatPermissions([]string{ "wfm:activityCode:add", "wfm:activityCode:delete", "wfm:activityCode:edit", "wfm:activityCode:view", "wfm:agent:edit", "wfm:agentSchedule:view", "wfm:agentTimeOffRequest:submit", "wfm:agent:view", "wfm:businessUnit:add", "wfm:businessUnit:delete", "wfm:businessUnit:edit", "wfm:businessUnit:view", "wfm:historicalAdherence:view", "wfm:shrinkage:view", "wfm:intraday:view", "wfm:managementUnit:add", "wfm:managementUnit:delete", "wfm:managementUnit:edit", "wfm:managementUnit:view", "wfm:publishedSchedule:view", "wfm:realtimeAdherence:view", "wfm:schedule:add", "wfm:schedule:delete", "wfm:schedule:edit", "wfm:schedule:generate", "wfm:schedule:view", "wfm:serviceGoalTemplate:add", "wfm:serviceGoalTemplate:delete", "wfm:serviceGoalTemplate:edit", "wfm:serviceGoalTemplate:view", "wfm:planningGroup:add", "wfm:planningGroup:delete", "wfm:planningGroup:edit", "wfm:planningGroup:view", "wfm:shiftTradeRequest:edit", "wfm:shiftTradeRequest:view", "wfm:agentShiftTradeRequest:participate", "wfm:shortTermForecast:add", "wfm:shortTermForecast:delete", "wfm:shortTermForecast:edit", "wfm:shortTermForecast:view", "wfm:staffingGroup:add", "wfm:staffingGroup:delete", "wfm:staffingGroup:edit", "wfm:staffingGroup:view", "wfm:timeOffRequest:add", "wfm:timeOffRequest:edit", "wfm:timeOffRequest:view", "wfm:timeOffLimit:add", "wfm:timeOffLimit:delete", "wfm:timeOffLimit:edit", "wfm:timeOffLimit:view", "wfm:timeOffPlan:add", "wfm:timeOffPlan:delete", "wfm:timeOffPlan:edit", "wfm:timeOffPlan:view", "wfm:timeOffRequest:add", "wfm:timeOffRequest:edit", "wfm:timeOffRequest:view", "wfm:workPlan:add", "wfm:workPlan:delete", "wfm:workPlan:edit", "wfm:workPlan:view", "wfm:workPlanRotation:add", "wfm:workPlanRotation:delete", "wfm:workPlanRotation:edit", "wfm:workPlanRotation:view", "coaching:appointment:add", "coaching:appointment:edit", "learning:assignment:add", "learning:assignment:reschedule",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BusinessUnitResponse"
      }
    }
  }
}`)
	workforcemanagement_businessunitsCmd.AddCommand(getCmd)

	utils.AddFlag(listCmd.Flags(), "string", "feature", "", "If specified, the list of business units for which the user is authorized to use the requested feature will be returned Valid values: AgentHistoricalAdherence, AgentHistoricalAdherenceConformance, AgentSchedule, AgentTimeOffRequest, AgentWorkPlanBid, AlternativeShift, Coaching, Learning, AgentUnavailableTimes, ActivityCodes, ActivityPlans, UnavailableTimes, Agents, BuActivityCodes, BusinessUnits, CapacityPlan, ContinuousForecast, HistoricalAdherence, HistoricalShrinkage, IntradayMonitoring, BuIntradayMonitoring, ManagementUnits, RealTimeAdherence, Schedules, BuSchedules, ServiceGoalTemplates, PlanningGroups, LongTermStaffing, ShiftTrading, ShortTermForecasts, BuShortTermForecasts, StaffingGroups, TimeOffPlans, TimeOffRequests, TimeOffLimits, WorkPlanBids, WorkPlanBidGroups, WorkPlanRotations, WorkPlans")
	utils.AddFlag(listCmd.Flags(), "string", "divisionId", "", "If specified, the list of business units belonging to the specified division will be returned")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BusinessUnitListing"
      }
    }
  }
}`)
	workforcemanagement_businessunitsCmd.AddCommand(listCmd)

	utils.AddFlag(updateCmd.Flags(), "bool", "includeSchedulingDefaultMessageSeverities", "", "Whether to include scheduling default message severities")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/workforcemanagement/businessunits/{businessUnitId}", utils.FormatPermissions([]string{ "wfm:businessUnit:edit",  }), utils.GenerateDevCentreLink("PATCH", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "description" : "body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UpdateBusinessUnitRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BusinessUnitResponse"
      }
    }
  }
}`)
	workforcemanagement_businessunitsCmd.AddCommand(updateCmd)
	return workforcemanagement_businessunitsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new business unit",
	Long:  "Add a new business unit",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createbusinessunitrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits"

		includeSchedulingDefaultMessageSeverities := utils.GetFlag(cmd.Flags(), "bool", "includeSchedulingDefaultMessageSeverities")
		if includeSchedulingDefaultMessageSeverities != "" {
			queryParams["includeSchedulingDefaultMessageSeverities"] = includeSchedulingDefaultMessageSeverities
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "create"
		const httpMethod = "POST"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [businessUnitId]",
	Short: "Delete business unit",
	Long:  "Delete business unit",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "delete"
		const httpMethod = "DELETE"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var getCmd = &cobra.Command{
	Use:   "get [businessUnitId]",
	Short: "Get business unit",
	Long:  "Get business unit",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)

		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
		includeSchedulingDefaultMessageSeverities := utils.GetFlag(cmd.Flags(), "bool", "includeSchedulingDefaultMessageSeverities")
		if includeSchedulingDefaultMessageSeverities != "" {
			queryParams["includeSchedulingDefaultMessageSeverities"] = includeSchedulingDefaultMessageSeverities
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "get"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get business units",
	Long:  "Get business units",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits"

		feature := utils.GetFlag(cmd.Flags(), "string", "feature")
		if feature != "" {
			queryParams["feature"] = feature
		}
		divisionId := utils.GetFlag(cmd.Flags(), "string", "divisionId")
		if divisionId != "" {
			queryParams["divisionId"] = divisionId
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "list"
		const httpMethod = "GET"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [businessUnitId]",
	Short: "Update business unit",
	Long:  "Update business unit",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Updatebusinessunitrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)

		includeSchedulingDefaultMessageSeverities := utils.GetFlag(cmd.Flags(), "bool", "includeSchedulingDefaultMessageSeverities")
		if includeSchedulingDefaultMessageSeverities != "" {
			queryParams["includeSchedulingDefaultMessageSeverities"] = includeSchedulingDefaultMessageSeverities
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
		}

		const opId = "update"
		const httpMethod = "PATCH"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
		// TODO read from config file
		retryConfig := &retry.RetryConfiguration{
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 60 * time.Second,
			RetryMax:     20,
		}
		results, err := retryFunc(retryConfig)
		if err != nil {
			if httpMethod == "HEAD" {
				if httpErr, ok := err.(models.HttpStatusError); ok {
					logger.Fatal(fmt.Sprintf("Status Code %v\n", httpErr.StatusCode))
				}
			}
			logger.Fatal(err)
		}

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
