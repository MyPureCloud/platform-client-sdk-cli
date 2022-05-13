package workforcemanagement_managementunits_workplans_copy

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
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_workplans_copy", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/copy", )
	workforcemanagement_managementunits_workplans_copyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_workplans_copy"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_workplans_copyCmd)
}

func Cmdworkforcemanagement_managementunits_workplans_copy() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/copy", utils.FormatPermissions([]string{ "wfm:workPlan:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/copy")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CopyWorkPlan"
      }
    }
  },
  "required" : false
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WorkPlan"
      }
    }
  }
}`)
	workforcemanagement_managementunits_workplans_copyCmd.AddCommand(createCmd)
	return workforcemanagement_managementunits_workplans_copyCmd
}

var createCmd = &cobra.Command{
	Use:   "create [managementUnitId] [workPlanId]",
	Short: "Create a copy of work plan",
	Long:  "Create a copy of work plan",
	Args:  utils.DetermineArgs([]string{ "managementUnitId", "workPlanId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Copyworkplan{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/copy"
		managementUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
		workPlanId, args := args[0], args[1:]
		path = strings.Replace(path, "{workPlanId}", fmt.Sprintf("%v", workPlanId), -1)

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
