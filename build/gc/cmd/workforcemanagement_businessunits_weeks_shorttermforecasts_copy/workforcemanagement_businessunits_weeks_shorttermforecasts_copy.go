package workforcemanagement_businessunits_weeks_shorttermforecasts_copy

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_shorttermforecasts_copy", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy", )
	workforcemanagement_businessunits_weeks_shorttermforecasts_copyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_shorttermforecasts_copy"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_weeks_shorttermforecasts_copyCmd)
}

func Cmdworkforcemanagement_businessunits_weeks_shorttermforecasts_copy() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "forceAsync", "", "Force the result of this operation to be sent asynchronously via notification.  For testing/app development purposes")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy", utils.FormatPermissions([]string{ "wfm:shortTermForecast:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "body",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/CopyBuForecastRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/AsyncForecastOperationResult"
  }
}`)
	workforcemanagement_businessunits_weeks_shorttermforecasts_copyCmd.AddCommand(createCmd)
	
	return workforcemanagement_businessunits_weeks_shorttermforecasts_copyCmd
}

var createCmd = &cobra.Command{
	Use:   "create [businessUnitId] [weekDateId] [forecastId]",
	Short: "Copy a short term forecast",
	Long:  "Copy a short term forecast",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "weekDateId", "forecastId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Copybuforecastrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
		weekDateId, args := args[0], args[1:]
		path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
		forecastId, args := args[0], args[1:]
		path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)

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
