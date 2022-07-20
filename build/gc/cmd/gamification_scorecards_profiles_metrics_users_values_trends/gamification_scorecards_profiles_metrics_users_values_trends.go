package gamification_scorecards_profiles_metrics_users_values_trends

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
	Description = utils.FormatUsageDescription("gamification_scorecards_profiles_metrics_users_values_trends", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/{userId}/values/trends", )
	gamification_scorecards_profiles_metrics_users_values_trendsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_profiles_metrics_users_values_trends"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_profiles_metrics_users_values_trendsCmd)
}

func Cmdgamification_scorecards_profiles_metrics_users_values_trends() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "string", "filterType", "", "Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division")
	utils.AddFlag(getCmd.Flags(), "string", "filterId", "", "ID for the filter type. For example, division Id - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "startWorkday", "", "Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "endWorkday", "", "End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "referenceWorkday", "", "Reference workday for the trend. Used to determine the associated metric definition. If not set, then the value of endWorkday is used. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	utils.AddFlag(getCmd.Flags(), "string", "timeZone", "UTC", "Timezone for the workday. Defaults to UTC")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends", utils.FormatPermissions([]string{ "gamification:scorecard:viewAll",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("filterType")
	getCmd.MarkFlagRequired("filterId")
	getCmd.MarkFlagRequired("startWorkday")
	getCmd.MarkFlagRequired("endWorkday")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/MetricValueTrendAverage"
      }
    }
  }
}`)
	gamification_scorecards_profiles_metrics_users_values_trendsCmd.AddCommand(getCmd)

	utils.AddFlag(getusertrendsCmd.Flags(), "time.Time", "startWorkday", "", "Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getusertrendsCmd.Flags(), "time.Time", "endWorkday", "", "End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getusertrendsCmd.Flags(), "time.Time", "referenceWorkday", "", "Reference workday for the trend. Used to determine the associated metric definition. If not set, then the value of endWorkday is used. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	utils.AddFlag(getusertrendsCmd.Flags(), "string", "timeZone", "UTC", "Timezone for the workday. Defaults to UTC")
	getusertrendsCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getusertrendsCmd.UsageTemplate(), "GET", "/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/{userId}/values/trends", utils.FormatPermissions([]string{ "gamification:scorecard:viewAll",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/{userId}/values/trends")))
	utils.AddFileFlagIfUpsert(getusertrendsCmd.Flags(), "GET", ``)
	getusertrendsCmd.MarkFlagRequired("startWorkday")
	getusertrendsCmd.MarkFlagRequired("endWorkday")
	
	utils.AddPaginateFlagsIfListingResponse(getusertrendsCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/MetricValueTrendAverage"
      }
    }
  }
}`)
	gamification_scorecards_profiles_metrics_users_values_trendsCmd.AddCommand(getusertrendsCmd)
	return gamification_scorecards_profiles_metrics_users_values_trendsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [profileId] [metricId]",
	Short: "Average performance values trends by metric of a division or a performance profile",
	Long:  "Average performance values trends by metric of a division or a performance profile",
	Args:  utils.DetermineArgs([]string{ "profileId", "metricId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends"
		profileId, args := args[0], args[1:]
		path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileId), -1)
		metricId, args := args[0], args[1:]
		path = strings.Replace(path, "{metricId}", fmt.Sprintf("%v", metricId), -1)

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
		referenceWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "referenceWorkday")
		if referenceWorkday != "" {
			queryParams["referenceWorkday"] = referenceWorkday
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
var getusertrendsCmd = &cobra.Command{
	Use:   "getusertrends [profileId] [metricId] [userId]",
	Short: "Average performance values trends by metric of a user",
	Long:  "Average performance values trends by metric of a user",
	Args:  utils.DetermineArgs([]string{ "profileId", "metricId", "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/{userId}/values/trends"
		profileId, args := args[0], args[1:]
		path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileId), -1)
		metricId, args := args[0], args[1:]
		path = strings.Replace(path, "{metricId}", fmt.Sprintf("%v", metricId), -1)
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
		referenceWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "referenceWorkday")
		if referenceWorkday != "" {
			queryParams["referenceWorkday"] = referenceWorkday
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

		const opId = "getusertrends"
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
