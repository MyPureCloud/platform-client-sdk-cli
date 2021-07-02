package webchat_guest_conversations_members_typing

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
	Description = utils.FormatUsageDescription("webchat_guest_conversations_members_typing", "SWAGGER_OVERRIDE_/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing", )
	webchat_guest_conversations_members_typingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webchat_guest_conversations_members_typing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webchat_guest_conversations_members_typingCmd)
}

func Cmdwebchat_guest_conversations_members_typing() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "WebChat", "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatTyping&quot;
  }
}`)
	webchat_guest_conversations_members_typingCmd.AddCommand(createCmd)
	
	return webchat_guest_conversations_members_typingCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId] [memberId]",
	Short: "Send a typing-indicator in a chat conversation.",
	Long:  "Send a typing-indicator in a chat conversation.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "memberId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing"
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
