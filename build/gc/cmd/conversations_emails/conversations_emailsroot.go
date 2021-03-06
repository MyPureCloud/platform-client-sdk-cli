package conversations_emails

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_emails_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_emails_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_emails_inboundmessages"
)

func init() {
	conversations_emailsCmd.AddCommand(conversations_emails_participants.Cmdconversations_emails_participants())
	conversations_emailsCmd.AddCommand(conversations_emails_messages.Cmdconversations_emails_messages())
	conversations_emailsCmd.AddCommand(conversations_emails_inboundmessages.Cmdconversations_emails_inboundmessages())
	conversations_emailsCmd.Short = utils.GenerateCustomDescription(conversations_emailsCmd.Short, conversations_emails_participants.Description, conversations_emails_messages.Description, conversations_emails_inboundmessages.Description, )
	conversations_emailsCmd.Long = conversations_emailsCmd.Short
}
