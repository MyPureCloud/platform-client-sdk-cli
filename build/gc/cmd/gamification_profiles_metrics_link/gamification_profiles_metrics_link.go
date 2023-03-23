package gamification_profiles_metrics_link

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
	Description = utils.FormatUsageDescription("gamification_profiles_metrics_link", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles/{sourceProfileId}/metrics/{sourceMetricId}/link", )
	gamification_profiles_metrics_linkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_profiles_metrics_link"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_profiles_metrics_linkCmd)
}

func Cmdgamification_profiles_metrics_link() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/gamification/profiles/{sourceProfileId}/metrics/{sourceMetricId}/link", utils.FormatPermissions([]string{ "gamification:profile:update",  }), utils.GenerateDevCentreLink("POST", "Gamification", "/api/v2/gamification/profiles/{sourceProfileId}/metrics/{sourceMetricId}/link")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "linkedMetric",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/TargetPerformanceProfile"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "Linked Metric successfully created",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/Metric"
      }
    }
  }
}`)
	gamification_profiles_metrics_linkCmd.AddCommand(createCmd)
	return gamification_profiles_metrics_linkCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [sourceProfileId] [sourceMetricId]",
	Short: "Creates a linked metric",
	Long:  "Creates a linked metric",
	Args:  utils.DetermineArgs([]string{ "sourceProfileId", "sourceMetricId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Targetperformanceprofile{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles/{sourceProfileId}/metrics/{sourceMetricId}/link"
		sourceProfileId, args := args[0], args[1:]
		path = strings.Replace(path, "{sourceProfileId}", fmt.Sprintf("%v", sourceProfileId), -1)
		sourceMetricId, args := args[0], args[1:]
		path = strings.Replace(path, "{sourceMetricId}", fmt.Sprintf("%v", sourceMetricId), -1)

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
