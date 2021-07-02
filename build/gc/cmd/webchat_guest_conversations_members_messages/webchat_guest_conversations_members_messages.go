package webchat_guest_conversations_members_messages

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
	Description = utils.FormatUsageDescription("webchat_guest_conversations_members_messages", "SWAGGER_OVERRIDE_/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages", )
	webchat_guest_conversations_members_messagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webchat_guest_conversations_members_messages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webchat_guest_conversations_members_messagesCmd)
}

func Cmdwebchat_guest_conversations_members_messages() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "WebChat", "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Message&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/CreateWebChatMessageRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatMessage&quot;
  }
}`)
	webchat_guest_conversations_members_messagesCmd.AddCommand(createCmd)
	
	return webchat_guest_conversations_members_messagesCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId] [memberId]",
	Short: "Send a message in a chat conversation.",
	Long:  "Send a message in a chat conversation.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		memberId, args := args[0], args[1:]
		path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)

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
