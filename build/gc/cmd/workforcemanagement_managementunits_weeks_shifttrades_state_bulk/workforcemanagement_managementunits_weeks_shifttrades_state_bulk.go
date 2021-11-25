package workforcemanagement_managementunits_weeks_shifttrades_state_bulk

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
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_weeks_shifttrades_state_bulk", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state/bulk", )
	workforcemanagement_managementunits_weeks_shifttrades_state_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_weeks_shifttrades_state_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_weeks_shifttrades_state_bulkCmd)
}

func Cmdworkforcemanagement_managementunits_weeks_shifttrades_state_bulk() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "forceAsync", "", "Force the result of this operation to be sent asynchronously via notification.  For testing/app development purposes")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state/bulk", utils.FormatPermissions([]string{ "wfm:agentShiftTradeRequest:participate", "wfm:shiftTradeRequest:edit",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state/bulk")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/BulkShiftTradeStateUpdateRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/BulkUpdateShiftTradeStateResponse"
  }
}`)
	workforcemanagement_managementunits_weeks_shifttrades_state_bulkCmd.AddCommand(createCmd)
	
	return workforcemanagement_managementunits_weeks_shifttrades_state_bulkCmd
}

var createCmd = &cobra.Command{
	Use:   "create [managementUnitId] [weekDateId]",
	Short: "Updates the state of a batch of shift trades",
	Long:  "Updates the state of a batch of shift trades",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "weekDateId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Bulkshifttradestateupdaterequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state/bulk"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		weekDateId, args := args[0], args[1:]
		path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)

		forceAsync := utils.GetFlag(cmd.Flags(), "bool", "forceAsync")
		if forceAsync != "" {
			queryParams["forceAsync"] = forceAsync
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
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
			logger.Fatal(err)
		}

		utils.Render(results)
	},
}
