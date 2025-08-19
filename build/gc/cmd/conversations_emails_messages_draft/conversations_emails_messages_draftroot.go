package conversations_emails_messages_draft

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_emails_messages_draft_attachments"
)

func init() {
	conversations_emails_messages_draftCmd.AddCommand(conversations_emails_messages_draft_attachments.Cmdconversations_emails_messages_draft_attachments())
	conversations_emails_messages_draftCmd.Short = utils.GenerateCustomDescription(conversations_emails_messages_draftCmd.Short, conversations_emails_messages_draft_attachments.Description, )
	conversations_emails_messages_draftCmd.Long = conversations_emails_messages_draftCmd.Short
}
