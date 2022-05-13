package analytics_reporting_schedules_runreport

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
	Description = utils.FormatUsageDescription("analytics_reporting_schedules_runreport", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting/schedules/{scheduleId}/runreport", )
	analytics_reporting_schedules_runreportCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting_schedules_runreport"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reporting_schedules_runreportCmd)
}

func Cmdanalytics_reporting_schedules_runreport() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/analytics/reporting/schedules/{scheduleId}/runreport", utils.FormatPermissions([]string{ "reporting:acd:view, reporting:status:view, reporting:interactions:view, reporting:outbound:view, reporting:quality:view, employee",  }), utils.GenerateDevCentreLink("POST", "Analytics", "/api/v2/analytics/reporting/schedules/{scheduleId}/runreport")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Accepted - Processing Report",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/RunNowResponse"
      }
    }
  }
}`)
	analytics_reporting_schedules_runreportCmd.AddCommand(createCmd)
	return analytics_reporting_schedules_runreportCmd
}

var createCmd = &cobra.Command{
	Use:   "create [scheduleId]",
	Short: "Place a scheduled report immediately into the reporting queue",
	Long:  "Place a scheduled report immediately into the reporting queue",
	Args:  utils.DetermineArgs([]string{ "scheduleId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/analytics/reporting/schedules/{scheduleId}/runreport"
		scheduleId, args := args[0], args[1:]
		path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)

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
