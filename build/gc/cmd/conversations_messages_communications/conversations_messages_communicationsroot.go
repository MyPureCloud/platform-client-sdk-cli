package conversations_messages_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_communications_messages"
)

func init() {
	conversations_messages_communicationsCmd.AddCommand(conversations_messages_communications_messages.Cmdconversations_messages_communications_messages())
	conversations_messages_communicationsCmd.Short = utils.GenerateCustomDescription(conversations_messages_communicationsCmd.Short, conversations_messages_communications_messages.Description, )
	conversations_messages_communicationsCmd.Long = conversations_messages_communicationsCmd.Short
}
