package conversations_messages_communications_messages_media

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
	Description = utils.FormatUsageDescription("conversations_messages_communications_messages_media", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media", )
	conversations_messages_communications_messages_mediaCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_communications_messages_media"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_communications_messages_mediaCmd)
}

func Cmdconversations_messages_communications_messages_media() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media", utils.FormatPermissions([]string{ "conversation:message:create", "conversation:webmessaging:create",  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", ``)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  &quot;description&quot; : &quot;Accepted&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/MessageMediaData&quot;
  }
}`)
	conversations_messages_communications_messages_mediaCmd.AddCommand(createCmd)
	
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media/{mediaId}", utils.FormatPermissions([]string{ "conversation:message:view", "conversation:webmessaging:view",  }), utils.GenerateDevCentreLink("GET", "Conversations", "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media/{mediaId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/MessageMediaData&quot;
  }
}`)
	conversations_messages_communications_messages_mediaCmd.AddCommand(getCmd)
	
	return conversations_messages_communications_messages_mediaCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId] [communicationId]",
	Short: "Create media",
	Long:  "Create media",
	Args:  utils.DetermineArgs([]string{ "conversationId", "communicationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media"
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
var getCmd = &cobra.Command{
	Use:   "get [conversationId] [communicationId] [mediaId]",
	Short: "Get media",
	Long:  "Get media",
	Args:  utils.DetermineArgs([]string{ "conversationId", "communicationId", "mediaId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media/{mediaId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		communicationId, args := args[0], args[1:]
		path = strings.Replace(path, "{communicationId}", fmt.Sprintf("%v", communicationId), -1)
		mediaId, args := args[0], args[1:]
		path = strings.Replace(path, "{mediaId}", fmt.Sprintf("%v", mediaId), -1)

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
