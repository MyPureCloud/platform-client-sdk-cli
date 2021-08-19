package gamification_profiles_metrics_objectivedetails

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
	Description = utils.FormatUsageDescription("gamification_profiles_metrics_objectivedetails", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles/{profileId}/metrics/objectivedetails", )
	gamification_profiles_metrics_objectivedetailsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_profiles_metrics_objectivedetails"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_profiles_metrics_objectivedetailsCmd)
}

func Cmdgamification_profiles_metrics_objectivedetails() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "time.Time", "workday", "", "The objective query workday. If not specified, then it retrieves the current objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/gamification/profiles/{profileId}/metrics/objectivedetails", utils.FormatPermissions([]string{ "gamification:profile:view", "gamification:leaderboard:view", "gamification:scorecard:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/profiles/{profileId}/metrics/objectivedetails")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/GetMetricsResponse&quot;
  }
}`)
	gamification_profiles_metrics_objectivedetailsCmd.AddCommand(listCmd)
	
	return gamification_profiles_metrics_objectivedetailsCmd
}

var listCmd = &cobra.Command{
	Use:   "list [profileId]",
	Short: "All metrics for a given performance profile with objective details such as order and maxPoints",
	Long:  "All metrics for a given performance profile with objective details such as order and maxPoints",
	Args:  utils.DetermineArgs([]string{ "profileId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles/{profileId}/metrics/objectivedetails"
		profileId, args := args[0], args[1:]
		path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileId), -1)

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
