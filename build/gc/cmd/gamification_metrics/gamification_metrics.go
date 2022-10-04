package gamification_metrics

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
	Description = utils.FormatUsageDescription("gamification_metrics", "SWAGGER_OVERRIDE_/api/v2/gamification/metrics", "SWAGGER_OVERRIDE_/api/v2/gamification/metrics", )
	gamification_metricsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_metrics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_metricsCmd)
}

func Cmdgamification_metrics() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "time.Time", "workday", "", "The objective query workday. If not specified, then it retrieves the current objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	utils.AddFlag(getCmd.Flags(), "string", "performanceProfileId", "", "The profile id of the metrics you are trying to retrieve. The DEFAULT profile is used if nothing is given.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/metrics/{metricId}", utils.FormatPermissions([]string{ "gamification:profile:view", "gamification:leaderboard:view", "gamification:scorecard:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/metrics/{metricId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Metric"
      }
    }
  }
}`)
	gamification_metricsCmd.AddCommand(getCmd)

	utils.AddFlag(updateCmd.Flags(), "string", "performanceProfileId", "", "The profile id of the metrics you are trying to retrieve. The DEFAULT profile is used if nothing is given.")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/gamification/metrics/{metricId}", utils.FormatPermissions([]string{ "gamification:profile:update",  }), utils.GenerateDevCentreLink("PUT", "Gamification", "/api/v2/gamification/metrics/{metricId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "description" : "Metric",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/CreateMetric"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Metric"
      }
    }
  }
}`)
	gamification_metricsCmd.AddCommand(updateCmd)
	return gamification_metricsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [metricId]",
	Short: "Gamified metric by id",
	Long:  "Gamified metric by id",
	Args:  utils.DetermineArgs([]string{ "metricId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/metrics/{metricId}"
		metricId, args := args[0], args[1:]
		path = strings.Replace(path, "{metricId}", fmt.Sprintf("%v", metricId), -1)

		workday := utils.GetFlag(cmd.Flags(), "time.Time", "workday")
		if workday != "" {
			queryParams["workday"] = workday
		}
		performanceProfileId := utils.GetFlag(cmd.Flags(), "string", "performanceProfileId")
		if performanceProfileId != "" {
			queryParams["performanceProfileId"] = performanceProfileId
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

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
var updateCmd = &cobra.Command{
	Use:   "update [metricId]",
	Short: "Updates a metric",
	Long:  "Updates a metric",
	Args:  utils.DetermineArgs([]string{ "metricId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Createmetric{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/metrics/{metricId}"
		metricId, args := args[0], args[1:]
		path = strings.Replace(path, "{metricId}", fmt.Sprintf("%v", metricId), -1)

		performanceProfileId := utils.GetFlag(cmd.Flags(), "string", "performanceProfileId")
		if performanceProfileId != "" {
			queryParams["performanceProfileId"] = performanceProfileId
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		const opId = "update"
		const httpMethod = "PUT"
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

		filterCondition, _ := cmd.Flags().GetString("filtercondition")
		if filterCondition != "" {
			filteredResults, err := utils.FilterByCondition(results, filterCondition)
			if err != nil {
				logger.Fatal(err)
			}
			results = filteredResults
		}

		utils.Render(results)
	},
}
