package gamification_leaderboard_bestpoints

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
	Description = utils.FormatUsageDescription("gamification_leaderboard_bestpoints", "SWAGGER_OVERRIDE_/api/v2/gamification/leaderboard/bestpoints", )
	gamification_leaderboard_bestpointsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_leaderboard_bestpoints"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_leaderboard_bestpointsCmd)
}

func Cmdgamification_leaderboard_bestpoints() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/gamification/leaderboard/bestpoints", utils.FormatPermissions([]string{ "gamification:leaderboard:view",  }), utils.GenerateDevCentreLink("GET", "Gamification", "/api/v2/gamification/leaderboard/bestpoints")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/OverallBestPoints&quot;
  }
}`)
	gamification_leaderboard_bestpointsCmd.AddCommand(getCmd)
	
	return gamification_leaderboard_bestpointsCmd
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Best Points of the requesting user`s current performance profile or division",
	Long:  "Best Points of the requesting user`s current performance profile or division",
	Args:  utils.DetermineArgs([]string{ }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/gamification/leaderboard/bestpoints"

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
