package gamification_profiles_users

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
	Description = utils.FormatUsageDescription("gamification_profiles_users", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles/users", )
	gamification_profiles_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_profiles_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_profiles_usersCmd)
}

func Cmdgamification_profiles_users() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "time.Time", "workday", "", "Target querying workday. If not provided, then queries the current performance profile. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/profiles/users/{userId}", utils.FormatPermissions([]string{ "gamification:profile:view", "gamification:scorecard:viewAll", "gamification:leaderboard:viewAll",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/profiles/users/{userId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/PerformanceProfile"
      }
    }
  }
}`)
	gamification_profiles_usersCmd.AddCommand(getCmd)
	return gamification_profiles_usersCmd
}

var getCmd = &cobra.Command{
	Use:   "get [userId]",
	Short: "Performance profile of a user",
	Long:  "Performance profile of a user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles/users/{userId}"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

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
