package journey_sessions

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
	Description = utils.FormatUsageDescription("journey_sessions", "SWAGGER_OVERRIDE_/api/v2/journey/sessions", )
	journey_sessionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_sessions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_sessionsCmd)
}

func Cmdjourney_sessions() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/journey/sessions/{sessionId}", utils.FormatPermissions([]string{ "journey:session:view",  }), utils.GenerateDevCentreLink("GET", "Journey", "/api/v2/journey/sessions/{sessionId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/Session&quot;
  }
}`)
	journey_sessionsCmd.AddCommand(getCmd)
	
	return journey_sessionsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [sessionId]",
	Short: "Retrieve a single session.",
	Long:  "Retrieve a single session.",
	Args:  utils.DetermineArgs([]string{ "sessionId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/journey/sessions/{sessionId}"
		sessionId, args := args[0], args[1:]
		path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)

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
