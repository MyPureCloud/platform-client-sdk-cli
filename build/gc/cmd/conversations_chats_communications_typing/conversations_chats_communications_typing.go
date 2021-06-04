package conversations_chats_communications_typing

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
	Description = utils.FormatUsageDescription("conversations_chats_communications_typing", "SWAGGER_OVERRIDE_/api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing", )
	conversations_chats_communications_typingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_chats_communications_typing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_chats_communications_typingCmd)
}

func Cmdconversations_chats_communications_typing() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing", utils.FormatPermissions([]string{  })))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatTyping&quot;
  }
}`)
	conversations_chats_communications_typingCmd.AddCommand(createCmd)
	
	return conversations_chats_communications_typingCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId] [communicationId]",
	Short: "Send a typing-indicator on behalf of a communication in a chat conversation.",
	Long:  "Send a typing-indicator on behalf of a communication in a chat conversation.",
	Args:  utils.DetermineArgs([]string{ "conversationId", "communicationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		communicationId, args := args[0], args[1:]
		path = strings.Replace(path, "{communicationId}", fmt.Sprintf("%v", communicationId), -1)

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
