package workforcemanagement_agents_me_managementunit

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
	Description = utils.FormatUsageDescription("workforcemanagement_agents_me_managementunit", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/me/managementunit", )
	workforcemanagement_agents_me_managementunitCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_me_managementunit"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_me_managementunitCmd)
}

func Cmdworkforcemanagement_agents_me_managementunit() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/agents/me/managementunit", utils.FormatPermissions([]string{ "wfm:agentSchedule:view", "wfm:agentTimeOffRequest:submit", "wfm:activityCode:add", "wfm:activityCode:delete", "wfm:activityCode:edit", "wfm:activityCode:view", "wfm:agent:edit", "wfm:agent:view", "wfm:businessUnit:add", "wfm:businessUnit:delete", "wfm:businessUnit:edit", "wfm:businessUnit:view", "wfm:historicalAdherence:view", "wfm:intraday:view", "wfm:managementUnit:add", "wfm:managementUnit:delete", "wfm:managementUnit:edit", "wfm:managementUnit:view", "wfm:realtimeAdherence:view", "wfm:schedule:add", "wfm:schedule:delete", "wfm:schedule:edit", "wfm:schedule:generate", "wfm:schedule:view", "wfm:publishedSchedule:view", "wfm:serviceGoalTemplate:add", "wfm:serviceGoalTemplate:delete", "wfm:serviceGoalTemplate:edit", "wfm:serviceGoalTemplate:view", "wfm:planningGroup:add", "wfm:planningGroup:delete", "wfm:planningGroup:edit", "wfm:planningGroup:view", "wfm:shiftTradeRequest:edit", "wfm:shiftTradeRequest:view", "wfm:shortTermForecast:add", "wfm:shortTermForecast:delete", "wfm:shortTermForecast:edit", "wfm:shortTermForecast:view", "wfm:timeOffLimit:add", "wfm:timeOffLimit:delete", "wfm:timeOffLimit:edit", "wfm:timeOffLimit:view", "wfm:timeOffPlan:add", "wfm:timeOffPlan:delete", "wfm:timeOffPlan:edit", "wfm:timeOffPlan:view", "wfm:timeOffRequest:add", "wfm:timeOffRequest:edit", "wfm:timeOffRequest:view", "wfm:workPlan:add", "wfm:workPlan:delete", "wfm:workPlan:edit", "wfm:workPlan:view", "wfm:workPlanRotation:add", "wfm:workPlanRotation:delete", "wfm:workPlanRotation:edit", "wfm:workPlanRotation:view",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/agents/me/managementunit")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/AgentManagementUnitReference"
      }
    }
  }
}`)
	workforcemanagement_agents_me_managementunitCmd.AddCommand(getCmd)
	return workforcemanagement_agents_me_managementunitCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the management unit to which the currently logged in agent belongs",
	Long:  "Get the management unit to which the currently logged in agent belongs",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/agents/me/managementunit"

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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

		utils.Render(results)
	},
}
