package conversations_messages_messages

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
	Description = utils.FormatUsageDescription("conversations_messages_messages", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/{conversationId}/messages", )
	conversations_messages_messagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_messages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_messagesCmd)
}

func Cmdconversations_messages_messages() *cobra.Command { 
	getCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", getCmd.UsageTemplate(), "GET", "/api/v2/conversations/messages/{conversationId}/messages/{messageId}", utils.FormatPermissions([]string{ "conversation:message:view", "conversation:webmessaging:view",  }), utils.GenerateDevCentreLink("GET", "Conversations", "/api/v2/conversations/messages/{conversationId}/messages/{messageId}")))
	utils.AddFileFlagIfUpsert(getCmd.Flags(), "GET", ``)
	
	utils.AddPaginateFlagsIfListingResponse(getCmd.Flags(), "GET", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/MessageData"
  }
}`)
	conversations_messages_messagesCmd.AddCommand(getCmd)
	
	return conversations_messages_messagesCmd
}

var getCmd = &cobra.Command{
	Use:   "get [conversationId] [messageId]",
	Short: "Get conversation message",
	Long:  "Get conversation message",
	Args:  utils.DetermineArgs([]string{ "conversationId", "messageId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/messages/{conversationId}/messages/{messageId}"
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
