package gamification_profiles_users_query

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
	Description = utils.FormatUsageDescription("gamification_profiles_users_query", "SWAGGER_OVERRIDE_/api/v2/gamification/profiles/users/{userId}/query", )
	gamification_profiles_users_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_profiles_users_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_profiles_users_queryCmd)
}

func Cmdgamification_profiles_users_query() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/gamification/profiles/users/{userId}/query", utils.FormatPermissions([]string{ "gamification:agentProfileMembership:view",  }), utils.GenerateDevCentreLink("POST", "Gamification", "/api/v2/gamification/profiles/users/{userId}/query")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "description" : "The date range of work day.",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UserProfilesInDateRangeRequest"
      }
    }
  },
  "required" : true
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/UserProfilesInDateRange"
      }
    }
  }
}`)
	gamification_profiles_users_queryCmd.AddCommand(createCmd)
	return gamification_profiles_users_queryCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var createCmd = &cobra.Command{
	Use:   "create [userId]",
	Short: "Query performance profiles in date range for a user",
	Long:  "Query performance profiles in date range for a user",
	Args:  utils.DetermineArgs([]string{ "userId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Userprofilesindaterangerequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/profiles/users/{userId}/query"
		userId, args := args[0], args[1:]
		path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)

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
