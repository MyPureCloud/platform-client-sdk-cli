package workforcemanagement_businessunits_scheduling_runs_result

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_scheduling_runs_result", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result", )
	workforcemanagement_businessunits_scheduling_runs_resultCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_scheduling_runs_result"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_scheduling_runs_resultCmd)
}

func Cmdworkforcemanagement_businessunits_scheduling_runs_result() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "[]string", "managementUnitIds", "", "The IDs of the management units for which to fetch the reschedule results - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "[]string", "expand", "", "The fields to expand. Omitting will return an empty response - REQUIRED Valid values: headcountForecast, generationResults, agentSchedules")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result", utils.FormatPermissions([]string{ "wfm:schedule:edit", "wfm:schedule:generate",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("managementUnitIds")
	getCmd.MarkFlagRequired("expand")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/BuRescheduleResult"
      }
    }
  }
}`)
	workforcemanagement_businessunits_scheduling_runs_resultCmd.AddCommand(getCmd)
	return workforcemanagement_businessunits_scheduling_runs_resultCmd
}

var getCmd = &cobra.Command{
	Use:   "get [businessUnitId] [runId]",
	Short: "Get the result of a rescheduling operation",
	Long:  "Get the result of a rescheduling operation",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "runId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
		runId, args := args[0], args[1:]
		path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)

		managementUnitIds := utils.GetFlag(cmd.Flags(), "[]string", "managementUnitIds")
		if managementUnitIds != "" {
			queryParams["managementUnitIds"] = managementUnitIds
		}
		expand := utils.GetFlag(cmd.Flags(), "[]string", "expand")
		if expand != "" {
			queryParams["expand"] = expand
		}
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
