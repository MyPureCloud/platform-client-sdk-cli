package gamification_scorecards_users_values_trends

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
	Description = utils.FormatUsageDescription("gamification_scorecards_users_values_trends", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/users/values/trends", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/users/{userId}/values/trends", )
	gamification_scorecards_users_values_trendsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_users_values_trends"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_users_values_trendsCmd)
}

func Cmdgamification_scorecards_users_values_trends() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "filterType", "", "Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division")
	utils.AddFlag(getCmd.Flags(), "string", "filterId", "", "ID for the filter type. - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "startWorkday", "", "Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "endWorkday", "", "End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "string", "timeZone", "UTC", "Timezone for the workday. Defaults to UTC")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/scorecards/users/values/trends", utils.FormatPermissions([]string{ "gamification:scorecard:viewAll",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/scorecards/users/values/trends")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("filterType")
	getCmd.MarkFlagRequired("filterId")
	getCmd.MarkFlagRequired("startWorkday")
	getCmd.MarkFlagRequired("endWorkday")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WorkdayValuesTrend&quot;
  }
}`)
	gamification_scorecards_users_values_trendsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "time.Time", "startWorkday", "", "Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "time.Time", "endWorkday", "", "End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "string", "timeZone", "UTC", "Timezone for the workday. Defaults to UTC")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/gamification/scorecards/users/{userId}/values/trends", utils.FormatPermissions([]string{ "gamification:scorecard:viewAll",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/scorecards/users/{userId}/values/trends")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("startWorkday")
	listCmd.MarkFlagRequired("endWorkday")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WorkdayValuesTrend&quot;
  }
}`)
	gamification_scorecards_users_values_trendsCmd.AddCommand(listCmd)
	
	return gamification_scorecards_users_values_trendsCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Values trend by target group",
	Long:  "Values trend by target group",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gamification/scorecards/users/values/trends"

		filterType := utils.GetFlag(cmd.Flags(), "string", "filterType")
		if filterType != "" {
			queryParams["filterType"] = filterType
		}
		filterId := utils.GetFlag(cmd.Flags(), "string", "filterId")
		if filterId != "" {
			queryParams["filterId"] = filterId
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

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
var listCmd = &cobra.Command{
	Use:   "list [userId]",
	Short: "Values Trends of a user",
	Long:  "Values Trends of a user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gamification/scorecards/users/{userId}/values/trends"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

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

		retryFunc := CommandService.DetermineAction("GET", urlString, cmd.Flags())
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
