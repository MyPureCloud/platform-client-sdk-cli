package workforcemanagement_businessunits_weeks_shorttermforecasts_import_uploadurl

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
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_shorttermforecasts_import_uploadurl", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/import/uploadurl", )
	workforcemanagement_businessunits_weeks_shorttermforecasts_import_uploadurlCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_shorttermforecasts_import_uploadurl"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_weeks_shorttermforecasts_import_uploadurlCmd)
}

func Cmdworkforcemanagement_businessunits_weeks_shorttermforecasts_import_uploadurl() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/import/uploadurl", utils.FormatPermissions([]string{ "wfm:shortTermForecast:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/import/uploadurl")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "body",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UploadUrlRequestBody"
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
        "$ref" : "#/components/schemas/ImportForecastUploadResponse"
      }
    }
  }
}`)
	workforcemanagement_businessunits_weeks_shorttermforecasts_import_uploadurlCmd.AddCommand(createCmd)
	return workforcemanagement_businessunits_weeks_shorttermforecasts_import_uploadurlCmd
}

var createCmd = &cobra.Command{
	Use:   "create [businessUnitId] [weekDateId]",
	Short: "Creates a signed upload URL for importing a short term forecast",
	Long:  "Creates a signed upload URL for importing a short term forecast",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "weekDateId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Uploadurlrequestbody{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/import/uploadurl"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
		weekDateId, args := args[0], args[1:]
		path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)

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
