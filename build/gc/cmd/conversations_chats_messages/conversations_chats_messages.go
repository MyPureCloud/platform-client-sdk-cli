package conversations_chats_messages

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
	Description = utils.FormatUsageDescription("conversations_chats_messages", "SWAGGER_OVERRIDE_/api/v2/conversations/chats/{conversationId}/messages", "SWAGGER_OVERRIDE_/api/v2/conversations/chats/{conversationId}/messages", )
	conversations_chats_messagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_chats_messages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_chats_messagesCmd)
}

func Cmdconversations_chats_messages() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/conversations/chats/{conversationId}/messages/{messageId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Conversations", "/api/v2/conversations/chats/{conversationId}/messages/{messageId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatMessage&quot;
  }
}`)
	conversations_chats_messagesCmd.AddCommand(getCmd)
	
	utils.AddFlag(listCmd.Flags(), "string", "after", "", "If specified, get the messages chronologically after the id of this message")
	utils.AddFlag(listCmd.Flags(), "string", "before", "", "If specified, get the messages chronologically before the id of this message")
	utils.AddFlag(listCmd.Flags(), "string", "sortOrder", "ascending", "Sort order Valid values: ascending, descending")
	utils.AddFlag(listCmd.Flags(), "int", "maxResults", "100", "Limit the returned number of messages, up to a maximum of 100")
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/conversations/chats/{conversationId}/messages", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "Conversations", "/api/v2/conversations/chats/{conversationId}/messages")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatMessageEntityList&quot;
  }
}`)
	conversations_chats_messagesCmd.AddCommand(listCmd)
	
	return conversations_chats_messagesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [conversationId] [messageId]",
	Short: "Get a web chat conversation message",
	Long:  "Get a web chat conversation message",
	Args:  utils.DetermineArgs([]string{ "conversationId", "messageId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/chats/{conversationId}/messages/{messageId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		messageId, args := args[0], args[1:]
		path = strings.Replace(path, "{messageId}", fmt.Sprintf("%v", messageId), -1)

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
var listCmd = &cobra.Command{
	Use:   "list [conversationId]",
	Short: "Get the messages of a chat conversation.",
	Long:  "Get the messages of a chat conversation.",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/chats/{conversationId}/messages"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

		after := utils.GetFlag(cmd.Flags(), "string", "after")
		if after != "" {
			queryParams["after"] = after
		}
		before := utils.GetFlag(cmd.Flags(), "string", "before")
		if before != "" {
			queryParams["before"] = before
		}
		sortOrder := utils.GetFlag(cmd.Flags(), "string", "sortOrder")
		if sortOrder != "" {
			queryParams["sortOrder"] = sortOrder
		}
		maxResults := utils.GetFlag(cmd.Flags(), "int", "maxResults")
		if maxResults != "" {
			queryParams["maxResults"] = maxResults
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
