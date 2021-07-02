package webchat_guest_conversations_mediarequests

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
	Description = utils.FormatUsageDescription("webchat_guest_conversations_mediarequests", "SWAGGER_OVERRIDE_/api/v2/webchat/guest/conversations/{conversationId}/mediarequests", "SWAGGER_OVERRIDE_/api/v2/webchat/guest/conversations/{conversationId}/mediarequests", "SWAGGER_OVERRIDE_/api/v2/webchat/guest/conversations/{conversationId}/mediarequests", )
	webchat_guest_conversations_mediarequestsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("webchat_guest_conversations_mediarequests"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(webchat_guest_conversations_mediarequestsCmd)
}

func Cmdwebchat_guest_conversations_mediarequests() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "WebChat", "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatGuestMediaRequest&quot;
  }
}`)
	webchat_guest_conversations_mediarequestsCmd.AddCommand(getCmd)
	
	listCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", listCmd.UsageTemplate(), "GET", "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("GET", "WebChat", "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests")))
	utils.AddFileFlagIfUpsert(listCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(listCmd.Flags(), "GET", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatGuestMediaRequestEntityList&quot;
  }
}`)
	webchat_guest_conversations_mediarequestsCmd.AddCommand(listCmd)
	
	updateCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", updateCmd.UsageTemplate(), "PATCH", "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("PATCH", "WebChat", "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}")))
	utils.AddFileFlagIfUpsert(updateCmd.Flags(), "PATCH", `{
  &quot;in&quot; : &quot;body&quot;,
  &quot;name&quot; : &quot;body&quot;,
  &quot;description&quot; : &quot;Request&quot;,
  &quot;required&quot; : true,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatGuestMediaRequest&quot;
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(updateCmd.Flags(), "PATCH", `{
  &quot;description&quot; : &quot;successful operation&quot;,
  &quot;schema&quot; : {
    &quot;$ref&quot; : &quot;#/definitions/WebChatGuestMediaRequest&quot;
  }
}`)
	webchat_guest_conversations_mediarequestsCmd.AddCommand(updateCmd)
	
	return webchat_guest_conversations_mediarequestsCmd
}

var getCmd = &cobra.Command{
	Use:   "get [conversationId] [mediaRequestId]",
	Short: "Get a media request in the conversation",
	Long:  "Get a media request in the conversation",
	Args:  utils.DetermineArgs([]string{ "conversationId", "mediaRequestId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		mediaRequestId, args := args[0], args[1:]
		path = strings.Replace(path, "{mediaRequestId}", fmt.Sprintf("%v", mediaRequestId), -1)

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
	Short: "Get all media requests to the guest in the conversation",
	Long:  "Get all media requests to the guest in the conversation",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)

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
var updateCmd = &cobra.Command{
	Use:   "update [conversationId] [mediaRequestId]",
	Short: "Update a media request in the conversation, setting the state to ACCEPTED/DECLINED/ERRORED",
	Long:  "Update a media request in the conversation, setting the state to ACCEPTED/DECLINED/ERRORED",
	Args:  utils.DetermineArgs([]string{ "conversationId", "mediaRequestId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		mediaRequestId, args := args[0], args[1:]
		path = strings.Replace(path, "{mediaRequestId}", fmt.Sprintf("%v", mediaRequestId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("PATCH", urlString, cmd.Flags())
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
