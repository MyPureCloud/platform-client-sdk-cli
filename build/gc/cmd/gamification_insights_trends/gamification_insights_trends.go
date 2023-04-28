package gamification_insights_trends

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
	Description = utils.FormatUsageDescription("gamification_insights_trends", "SWAGGER_OVERRIDE_/api/v2/gamification/insights/trends", )
	gamification_insights_trendsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_insights_trends"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_insights_trendsCmd)
}

func Cmdgamification_insights_trends() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "string", "filterType", "", "Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division")
	utils.AddFlag(listCmd.Flags(), "string", "filterId", "", "ID for the filter type. - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "string", "granularity", "", "Granularity - REQUIRED Valid values: Daily, Weekly")
	utils.AddFlag(listCmd.Flags(), "time.Time", "comparativePeriodStartWorkday", "", "The start work day of comparative period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "time.Time", "comparativePeriodEndWorkday", "", "The end work day of comparative period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "time.Time", "primaryPeriodStartWorkday", "", "The start work day of primary period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(listCmd.Flags(), "time.Time", "primaryPeriodEndWorkday", "", "The end work day of primary period. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/gamification/insights/trends", utils.FormatPermissions([]string{ "gamification:insights:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/insights/trends")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	listCmd.MarkFlagRequired("filterType")
	listCmd.MarkFlagRequired("filterId")
	listCmd.MarkFlagRequired("granularity")
	listCmd.MarkFlagRequired("comparativePeriodStartWorkday")
	listCmd.MarkFlagRequired("comparativePeriodEndWorkday")
	listCmd.MarkFlagRequired("primaryPeriodStartWorkday")
	listCmd.MarkFlagRequired("primaryPeriodEndWorkday")
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UserInsightsTrend"
      }
    }
  }
}`)
	gamification_insights_trendsCmd.AddCommand(listCmd)
	return gamification_insights_trendsCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get insights user trend for the current user",
	Long:  "Get insights user trend for the current user",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/insights/trends"

		filterType := utils.GetFlag(cmd.Flags(), "string", "filterType")
		if filterType != "" {
			queryParams["filterType"] = filterType
		}
		filterId := utils.GetFlag(cmd.Flags(), "string", "filterId")
		if filterId != "" {
			queryParams["filterId"] = filterId
		}
		granularity := utils.GetFlag(cmd.Flags(), "string", "granularity")
		if granularity != "" {
			queryParams["granularity"] = granularity
		}
		comparativePeriodStartWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "comparativePeriodStartWorkday")
		if comparativePeriodStartWorkday != "" {
			queryParams["comparativePeriodStartWorkday"] = comparativePeriodStartWorkday
		}
		comparativePeriodEndWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "comparativePeriodEndWorkday")
		if comparativePeriodEndWorkday != "" {
			queryParams["comparativePeriodEndWorkday"] = comparativePeriodEndWorkday
		}
		primaryPeriodStartWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "primaryPeriodStartWorkday")
		if primaryPeriodStartWorkday != "" {
			queryParams["primaryPeriodStartWorkday"] = primaryPeriodStartWorkday
		}
		primaryPeriodEndWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "primaryPeriodEndWorkday")
		if primaryPeriodEndWorkday != "" {
			queryParams["primaryPeriodEndWorkday"] = primaryPeriodEndWorkday
		}
		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", queryEscape(strings.TrimSpace(k)), queryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		if strings.Contains(urlString, "varType") {
			urlString = strings.Replace(urlString, "varType", "type", -1)
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
