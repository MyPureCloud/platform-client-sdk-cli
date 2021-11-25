package conversations_emails_messages_draft_attachments_copy

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
	Description = utils.FormatUsageDescription("conversations_emails_messages_draft_attachments_copy", "SWAGGER_OVERRIDE_/api/v2/conversations/emails/{conversationId}/messages/draft/attachments/copy", )
	conversations_emails_messages_draft_attachments_copyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_emails_messages_draft_attachments_copy"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_emails_messages_draft_attachments_copyCmd)
}

func Cmdconversations_emails_messages_draft_attachments_copy() *cobra.Command { 
	createCmd.SetUsageTemplate(fmt.Sprintf("%s\nOperation:\n  %s %s\n%s\n%s", createCmd.UsageTemplate(), "POST", "/api/v2/conversations/emails/{conversationId}/messages/draft/attachments/copy", utils.FormatPermissions([]string{  }), utils.GenerateDevCentreLink("POST", "Conversations", "/api/v2/conversations/emails/{conversationId}/messages/draft/attachments/copy")))
	utils.AddFileFlagIfUpsert(createCmd.Flags(), "POST", `{
  "in" : "body",
  "name" : "body",
  "description" : "Copy Attachment Request",
  "required" : true,
  "schema" : {
    "$ref" : "#/definitions/CopyAttachmentsRequest"
  }
}`)
	
	utils.AddPaginateFlagsIfListingResponse(createCmd.Flags(), "POST", `{
  "description" : "successful operation",
  "schema" : {
    "$ref" : "#/definitions/EmailMessage"
  }
}`)
	conversations_emails_messages_draft_attachments_copyCmd.AddCommand(createCmd)
	
	return conversations_emails_messages_draft_attachments_copyCmd
}

var createCmd = &cobra.Command{
	Use:   "create [conversationId]",
	Short: "Copy attachments from an email message to the current draft.",
	Long:  "Copy attachments from an email message to the current draft.",
	Args:  utils.DetermineArgs([]string{ "conversationId", }),

	Run: func(cmd *cobra.Command, args []string) {
		_ = models.Entities{}

		printReqBody, _ := cmd.Flags().GetBool("printrequestbody")
		if printReqBody {
			
			reqModel := models.Copyattachmentsrequest{}
			utils.Render(reqModel.String())
			
			return
		}

		queryParams := make(map[string]string)

		path := "/api/v2/conversations/emails/{conversationId}/messages/draft/attachments/copy"
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

		const opId = "create"
		const httpMethod = "POST"
		retryFunc := CommandService.DetermineAction(httpMethod, urlString, cmd, opId)
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
