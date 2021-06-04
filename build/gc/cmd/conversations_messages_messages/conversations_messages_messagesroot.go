package conversations_messages_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_messages_bulk"
)

func init() {
	conversations_messages_messagesCmd.AddCommand(conversations_messages_messages_bulk.Cmdconversations_messages_messages_bulk())
	conversations_messages_messagesCmd.Short = utils.GenerateCustomDescription(conversations_messages_messagesCmd.Short, conversations_messages_messages_bulk.Description, )
	conversations_messages_messagesCmd.Long = conversations_messages_messagesCmd.Short
}
