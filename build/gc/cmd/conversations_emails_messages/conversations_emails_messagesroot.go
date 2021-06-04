package conversations_emails_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_emails_messages_draft"
)

func init() {
	conversations_emails_messagesCmd.AddCommand(conversations_emails_messages_draft.Cmdconversations_emails_messages_draft())
	conversations_emails_messagesCmd.Short = utils.GenerateCustomDescription(conversations_emails_messagesCmd.Short, conversations_emails_messages_draft.Description, )
	conversations_emails_messagesCmd.Long = conversations_emails_messagesCmd.Short
}
