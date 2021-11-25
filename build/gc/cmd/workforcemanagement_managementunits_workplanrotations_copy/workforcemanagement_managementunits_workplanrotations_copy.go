package workforcemanagement_managementunits_workplanrotations_copy

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
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_workplanrotations_copy", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy", )
	workforcemanagement_managementunits_workplanrotations_copyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_workplanrotations_copy"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_workplanrotations_copyCmd)
}

func Cmdworkforcemanagement_managementunits_workplanrotations_copy() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy", utils.FormatPermissions([]string{ "wfm:workPlanRotation:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : false,
  "schema" : {
    "$ref" : "#/definitions/CopyWorkPlanRotationRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/WorkPlanRotationResponse"
  }
}`)
	workforcemanagement_managementunits_workplanrotations_copyCmd.AddCommand(createCmd)
	
	return workforcemanagement_managementunits_workplanrotations_copyCmd
}

var createCmd = &cobra.Command{
	Use:   "create [managementUnitId] [workPlanRotationId]",
	Short: "Create a copy of work plan rotation",
	Long:  "Create a copy of work plan rotation",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "workPlanRotationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Copyworkplanrotationrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		workPlanRotationId, args := args[0], args[1:]
		path = strings.Replace(path, "{workPlanRotationId}", fmt.Sprintf("%v", workPlanRotationId), -1)

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
