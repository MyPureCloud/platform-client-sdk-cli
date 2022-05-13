package workforcemanagement_businessunits_managementunits

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_managementunits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits", )
	workforcemanagement_businessunits_managementunitsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_managementunits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_managementunitsCmd)
}

func Cmdworkforcemanagement_businessunits_managementunits() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "feature", "", " Valid values: AgentSchedule, AgentTimeOffRequest, Coaching, CoachingDivisioned, ActivityCodes, Agents, BuActivityCodes, BusinessUnits, HistoricalAdherence, HistoricalShrinkage, IntradayMonitoring, BuIntradayMonitoring, ManagementUnits, RealTimeAdherence, Schedules, BuSchedules, ServiceGoalTemplates, PlanningGroups, ShiftTrading, ShortTermForecasts, BuShortTermForecasts, TimeOffPlans, TimeOffRequests, TimeOffLimits, WorkPlanRotations, WorkPlans")
	utils.AddFlag(listCmd.Flags(), "string", "divisionId", "", "")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	workforcemanagement_businessunits_managementunitsCmd.AddCommand(listCmd)
	return workforcemanagement_businessunits_managementunitsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [businessUnitId]",
	Short: "Get all authorized management units in the business unit",
	Long:  "Get all authorized management units in the business unit",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)

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
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
