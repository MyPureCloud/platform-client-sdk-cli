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
	Description = utils.FormatUsageDescription("gamification_metrics", "SWAGGER_OVERRIDE_/api/v2/gamification/metrics", "SWAGGER_OVERRIDE_/api/v2/gamification/metrics", "SWAGGER_OVERRIDE_/api/v2/gamification/metrics", "SWAGGER_OVERRIDE_/api/v2/gamification/metrics", )
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
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/gamification/metrics", utils.FormatPermissions([]string{ "gamification:profile:update",  }), utils.GenerateDevCentreLink("POST", "Gamification", "/api/v2/gamification/metrics")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Metric",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/Metric"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Metric successfully created",
  "schema" : {
    "$ref" : "#/definitions/Metric"
  }
}`)
	gamification_metricsCmd.AddCommand(createCmd)
	
	utils.AddFlag(getCmd.Flags(), "time.Time", "workday", "", "The objective query workday. If not specified, then it retrieves the current objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	utils.AddFlag(getCmd.Flags(), "string", "performanceProfileId", "", "The profile id of the metrics you are trying to retrieve. The DEFAULT profile is used if nothing is given.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/metrics/{metricId}", utils.FormatPermissions([]string{ "gamification:profile:view", "gamification:leaderboard:view", "gamification:scorecard:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/metrics/{metricId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Metric"
  }
}`)
	gamification_metricsCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "performanceProfileId", "", "The profile id of the metrics you are trying to retrieve. The DEFAULT profile is used if nothing is given.")
	utils.AddFlag(listCmd.Flags(), "time.Time", "workday", "", "The objective query workday. If not specified, then it retrieves the current objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/gamification/metrics", utils.FormatPermissions([]string{ "gamification:profile:view", "gamification:leaderboard:view", "gamification:scorecard:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/metrics")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/GetMetricsResponse"
  }
}`)
	gamification_metricsCmd.AddCommand(listCmd)
	
	utils.AddFlag(updateCmd.Flags(), "string", "performanceProfileId", "", "The profile id of the metrics you are trying to retrieve. The DEFAULT profile is used if nothing is given.")
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PUT", "/api/v2/gamification/metrics/{metricId}", utils.FormatPermissions([]string{ "gamification:profile:update",  }), utils.GenerateDevCentreLink("PUT", "Gamification", "/api/v2/gamification/metrics/{metricId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PUT", `{
  "in" : "body",
  "name" : "body",
  "description" : "Metric",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/Metric"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PUT", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Metric"
  }
}`)
	gamification_metricsCmd.AddCommand(updateCmd)
	
	return gamification_metricsCmd
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a gamified metric with a given metric definition and metric objective",
	Long:  "Creates a gamified metric with a given metric definition and metric objective",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Metric{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/metrics"


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
	Use:   "list",
	Short: "All gamified metrics for a given profile",
	Long:  "All gamified metrics for a given profile",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/metrics"


		performanceProfileId := utils.GetFlag(cmd.Flags(), "string", "performanceProfileId")
		if performanceProfileId != "" {
			queryParams["performanceProfileId"] = performanceProfileId
		}
		workday := utils.GetFlag(cmd.Flags(), "time.Time", "workday")
		if workday != "" {
			queryParams["workday"] = workday
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
var updateCmd = &cobra.Command{
	Use:   "update [metricId]",
	Short: "Updates a metric",
	Long:  "Updates a metric",
	Args:  utils.DetermineArgs([]string{ "metricId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Metric{}
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

		retryFunc := CommandService.DetermineAction("PUT", urlString, cmd.Flags())
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
