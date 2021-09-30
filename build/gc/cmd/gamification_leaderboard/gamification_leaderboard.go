package gamification_leaderboard

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
	Description = utils.FormatUsageDescription("gamification_leaderboard", "SWAGGER_OVERRIDE_/api/v2/gamification/leaderboard", )
	gamification_leaderboardCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_leaderboard"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_leaderboardCmd)
}

func Cmdgamification_leaderboard() *cobra.Command { 
	utils.AddFlag(getCmd.Flags(), "time.Time", "startWorkday", "", "Start workday to retrieve for the leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "time.Time", "endWorkday", "", "End workday to retrieve for the leaderboard. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED")
	utils.AddFlag(getCmd.Flags(), "string", "metricId", "", "Metric Id for which the leaderboard is to be generated. The total points is used if nothing is given.")
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/leaderboard", utils.FormatPermissions([]string{ "gamification:leaderboard:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/leaderboard")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	getCmd.MarkFlagRequired("startWorkday")
	getCmd.MarkFlagRequired("endWorkday")
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/Leaderboard"
  }
}`)
	gamification_leaderboardCmd.AddCommand(getCmd)
	
	return gamification_leaderboardCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Leaderboard of the requesting user`s division or performance profile",
	Long:  "Leaderboard of the requesting user`s division or performance profile",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/gamification/leaderboard"


		startWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "startWorkday")
		if startWorkday != "" {
			queryParams["startWorkday"] = startWorkday
		}
		endWorkday := utils.GetFlag(cmd.Flags(), "time.Time", "endWorkday")
		if endWorkday != "" {
			queryParams["endWorkday"] = endWorkday
		}
		metricId := utils.GetFlag(cmd.Flags(), "string", "metricId")
		if metricId != "" {
			queryParams["metricId"] = metricId
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
