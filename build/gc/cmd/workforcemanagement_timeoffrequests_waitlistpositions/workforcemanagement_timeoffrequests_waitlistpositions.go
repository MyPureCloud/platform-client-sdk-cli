package workforcemanagement_timeoffrequests_waitlistpositions

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
	Description = utils.FormatUsageDescription("workforcemanagement_timeoffrequests_waitlistpositions", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}/waitlistpositions", )
	workforcemanagement_timeoffrequests_waitlistpositionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_timeoffrequests_waitlistpositions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_timeoffrequests_waitlistpositionsCmd)
}

func Cmdworkforcemanagement_timeoffrequests_waitlistpositions() *cobra.Command { 
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}/waitlistpositions", utils.FormatPermissions([]string{ "wfm:agentTimeOffRequest:submit",  }), utils.GenerateDevCentreLink("GET", "Workforce Management", "/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}/waitlistpositions")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/WaitlistPositionListing"
      }
    }
  }
}`)
	workforcemanagement_timeoffrequests_waitlistpositionsCmd.AddCommand(listCmd)
	return workforcemanagement_timeoffrequests_waitlistpositionsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [timeOffRequestId]",
	Short: "Get the daily waitlist positions of a time off request for the current user",
	Long:  "Get the daily waitlist positions of a time off request for the current user",
	Args:  utils.DetermineArgs([]string{ "timeOffRequestId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}/waitlistpositions"
		timeOffRequestId, args := args[0], args[1:]
		path = strings.Replace(path, "{timeOffRequestId}", fmt.Sprintf("%v", timeOffRequestId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "list"
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
