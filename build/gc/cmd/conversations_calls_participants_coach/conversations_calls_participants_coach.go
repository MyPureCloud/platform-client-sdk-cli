package conversations_calls_participants_coach

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
	Description = utils.FormatUsageDescription("conversations_calls_participants_coach", "SWAGGER_OVERRIDE_/api/v2/conversations/calls/{conversationId}/participants/{participantId}/coach", )
	conversations_calls_participants_coachCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_calls_participants_coach"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_calls_participants_coachCmd)
}

func Cmdconversations_calls_participants_coach() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/calls/{conversationId}/participants/{participantId}/coach", utils.FormatPermissions([]string{ "conversation:call:coach",  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/calls/{conversationId}/participants/{participantId}/coach")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Created&quot;
}`)
	conversations_calls_participants_coachCmd.AddCommand(createCmd)
	
	return conversations_calls_participants_coachCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId] [participantId]",
	Short: "Listen in on the conversation from the point of view of a given participant while speaking to just the given participant.",
	Long:  "Listen in on the conversation from the point of view of a given participant while speaking to just the given participant.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "participantId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/calls/{conversationId}/participants/{participantId}/coach"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		participantId, args := args[0], args[1:]
		path = strings.Replace(path, "{participantId}", fmt.Sprintf("%v", participantId), -1)

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
