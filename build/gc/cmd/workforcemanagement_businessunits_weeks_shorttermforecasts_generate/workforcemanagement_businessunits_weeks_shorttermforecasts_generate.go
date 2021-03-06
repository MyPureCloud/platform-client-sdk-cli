package workforcemanagement_businessunits_weeks_shorttermforecasts_generate

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/retry"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
	"time"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_shorttermforecasts_generate", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate", )
	workforcemanagement_businessunits_weeks_shorttermforecasts_generateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_shorttermforecasts_generate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_weeks_shorttermforecasts_generateCmd)
}

func Cmdworkforcemanagement_businessunits_weeks_shorttermforecasts_generate() *cobra.Command { 
	utils.AddFlag(createCmd.Flags(), "bool", "forceAsync", "", "Force the result of this operation to be sent asynchronously via notification.  For testing/app development purposes")
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate", utils.FormatPermissions([]string{ "wfm:shortTermForecast:add",  }), utils.GenerateDevCentreLink("POST", "Workforce Management", "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/GenerateBuForecastRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/AsyncForecastOperationResult&quot;
  }
}`)
	workforcemanagement_businessunits_weeks_shorttermforecasts_generateCmd.AddCommand(createCmd)
	
	return workforcemanagement_businessunits_weeks_shorttermforecasts_generateCmd
}

var createCmd = &cobra.Command{
	Use:   "create [businessUnitId] [weekDateId]",
	Short: "Generate a short term forecast",
	Long:  "Generate a short term forecast",
	Args:  utils.DetermineArgs([]string{ "businessUnitId", "weekDateId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate"
		businessUnitId, args := args[0], args[1:]
		path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
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

		retryFunc := CommandService.DetermineAction("POST", urlString, cmd.Flags())
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
