package gamification_contests_agents_scores

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
	Description = utils.FormatUsageDescription("gamification_contests_agents_scores", "SWAGGER_OVERRIDE_/api/v2/gamification/contests/{contestId}/agents/scores", )
	gamification_contests_agents_scoresCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_contests_agents_scores"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_contests_agents_scoresCmd)
}

func Cmdgamification_contests_agents_scores() *cobra.Command { 
	utils.AddFlag(listCmd.Flags(), "int", "pageNumber", "1", "")
	utils.AddFlag(listCmd.Flags(), "int", "pageSize", "25", "")
	utils.AddFlag(listCmd.Flags(), "time.Time", "workday", "", "Target querying workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd")
	utils.AddFlag(listCmd.Flags(), "string", "returnsView", "All", "Desired response results Valid values: All, TopAndBottom")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/gamification/contests/{contestId}/agents/scores", utils.FormatPermissions([]string{ "gamification:contest:viewAll",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/contests/{contestId}/agents/scores")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "content" : {
    "application/json" : {
      "schema" : {
        "$ref" : "#/components/schemas/SWAGGER_OVERRIDE_list"
      }
    }
  }
}`)
	gamification_contests_agents_scoresCmd.AddCommand(listCmd)
	return gamification_contests_agents_scoresCmd
}

/* function introduced to differentiate string named 'url' from some service queryParams and /net/url imports */
func queryEscape(value string) string {
   return url.QueryEscape(value)
}

var listCmd = &cobra.Command{
	Use:   "list [contestId]",
	Short: "Get Contest Scores (Admin)",
	Long:  "Get Contest Scores (Admin)",
	Args:  utils.DetermineArgs([]string{ "contestId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/contests/{contestId}/agents/scores"
		contestId, args := args[0], args[1:]
		path = strings.Replace(path, "{contestId}", fmt.Sprintf("%v", contestId), -1)

		pageNumber := utils.GetFlag(cmd.Flags(), "int", "pageNumber")
		if pageNumber != "" {
			queryParams["pageNumber"] = pageNumber
		}
		pageSize := utils.GetFlag(cmd.Flags(), "int", "pageSize")
		if pageSize != "" {
			queryParams["pageSize"] = pageSize
		}
		workday := utils.GetFlag(cmd.Flags(), "time.Time", "workday")
		if workday != "" {
			queryParams["workday"] = workday
		}
		returnsView := utils.GetFlag(cmd.Flags(), "string", "returnsView")
		if returnsView != "" {
			queryParams["returnsView"] = returnsView
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
