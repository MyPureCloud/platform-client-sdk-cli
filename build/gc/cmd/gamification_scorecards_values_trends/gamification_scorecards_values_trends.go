package gamification_scorecards_values_trends

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
	Description = utils.FormatUsageDescription("gamification_scorecards_values_trends", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/values/trends", )
	gamification_scorecards_values_trendsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_values_trends"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_values_trendsCmd)
}

func Cmdgamification_scorecards_values_trends() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "filterType", "", "Filter type for the query request. If not set, then the request is for the requesting user. Valid values: PerformanceProfile, Division")
	utils.AddFlag(getCmd.Flags(), "time.Time", "referenceWorkday", "", "Reference workday for the trend. Used to determine the profile of the user as of this date. If not set, then the user`s current profile will be used. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	utils.AddFlag(getCmd.Flags(), "time.Time", "startWorkday", "", "Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "endWorkday", "", "End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "string", "timeZone", "UTC", "Timezone for the workday. Defaults to UTC")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/scorecards/values/trends", utils.FormatPermissions([]string{ "gamification:scorecard:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/scorecards/values/trends")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("startWorkday")
	getCmd.MarkFlagRequired("endWorkday")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/WorkdayValuesTrend"
  }
}`)
	gamification_scorecards_values_trendsCmd.AddCommand(getCmd)
	
	return gamification_scorecards_values_trendsCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Values trends of the requesting user or group",
	Long:  "Values trends of the requesting user or group",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/scorecards/values/trends"

		filterType := utils.GetFlag(cmd.Flags(), "string", "filterType")
		if filterType != "" {
			queryParams["filterType"] = filterType
		}
		referenceWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "referenceWorkday")
		if referenceWorkday != "" {
			queryParams["referenceWorkday"] = referenceWorkday
		}
		startWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "startWorkday")
		if startWorkday != "" {
			queryParams["startWorkday"] = startWorkday
		}
		endWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "endWorkday")
		if endWorkday != "" {
			queryParams["endWorkday"] = endWorkday
		}
		timeZone := utils.GetFlag(cmd.Flags(), "string", "timeZone")
		if timeZone != "" {
			queryParams["timeZone"] = timeZone
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
