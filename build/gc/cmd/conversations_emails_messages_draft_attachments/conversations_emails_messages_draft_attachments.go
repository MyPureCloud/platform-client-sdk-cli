package conversations_emails_messages_draft_attachments

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
	Description = utils.FormatUsageDescription("conversations_emails_messages_draft_attachments", "SWAGGER_OVERRIDE_/api/v2/conversations/emails/{conversationId}/messages/draft/attachments", )
	conversations_emails_messages_draft_attachmentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_emails_messages_draft_attachments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_emails_messages_draft_attachmentsCmd)
}

func Cmdconversations_emails_messages_draft_attachments() *cobra.Command { 
	deleteCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", deleteCmd.UsageTemplate(), "DELETE", "/api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("DELETE", "Conversations", "/api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}")))
	utils.AddFileFlagIfUpsert(deleteCmd.Flags(), "DELETE", ``)
	
	utils.AddPaginateFlagsIfListingResponse(deleteCmd.Flags(), "DELETE", `{
  &quot;description&quot; : &quot;Operation was successful.&quot;
}`)
	conversations_emails_messages_draft_attachmentsCmd.AddCommand(deleteCmd)
	
	return conversations_emails_messages_draft_attachmentsCmd
}

var deleteCmd = &cobra.Command{
	Use:   "delete [conversationId] [attachmentId]",
	Short: "Delete attachment from draft",
	Long:  "Delete attachment from draft",
	Args:  utils.DetermineArgs([]string{ "conversationId", "attachmentId", }),

	Run: func(cmd *cobra.Command, args []string) {
		queryParams := make(map[string]string)

		path := "/api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}"
		conversationId, args := args[0], args[1:]
		path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
		attachmentId, args := args[0], args[1:]
		path = strings.Replace(path, "{attachmentId}", fmt.Sprintf("%v", attachmentId), -1)

		urlString := path
		if len(queryParams) > 0 {
			urlString = fmt.Sprintf("%v?", path)
			for k, v := range queryParams {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
			urlString = strings.TrimSuffix(urlString, "&")
		}

		retryFunc := CommandService.DetermineAction("DELETE", urlString, cmd.Flags())
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
