package conversations_chats_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_communications_typing"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_communications_messages"
)

func init() {
	conversations_chats_communicationsCmd.AddCommand(conversations_chats_communications_typing.Cmdconversations_chats_communications_typing())
	conversations_chats_communicationsCmd.AddCommand(conversations_chats_communications_messages.Cmdconversations_chats_communications_messages())
	conversations_chats_communicationsCmd.Short = utils.GenerateCustomDescription(conversations_chats_communicationsCmd.Short, conversations_chats_communications_typing.Description, conversations_chats_communications_messages.Description, )
	conversations_chats_communicationsCmd.Long = conversations_chats_communicationsCmd.Short
}
