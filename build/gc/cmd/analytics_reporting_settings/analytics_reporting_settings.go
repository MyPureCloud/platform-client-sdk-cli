package analytics_reporting_settings

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
	Description = utils.FormatUsageDescription("analytics_reporting_settings", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting/settings", )
	analytics_reporting_settingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting_settings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reporting_settingsCmd)
}

func Cmdanalytics_reporting_settings() *cobra.Command { 
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/analytics/reporting/settings", utils.FormatPermissions([]string{ "recording:recordingSegment:view", "analytics:conversationDetail:view", "analytics:dashboardConfigurations:view",  }), utils.GenerateDevCentreLink("PATCH", "Analytics", "/api/v2/analytics/reporting/settings")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  "in" : "body",
  "name" : "body",
  "description" : "AnalyticsReportingSettingsRequest",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/AnalyticsReportingSettings"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/AnalyticsReportingSettings"
  }
}`)
	analytics_reporting_settingsCmd.AddCommand(updateCmd)
	
	return analytics_reporting_settingsCmd
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Patch AnalyticsReportingSettings values for an organization",
	Long:  "Patch AnalyticsReportingSettings values for an organization",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Analyticsreportingsettings{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/reporting/settings"


		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
