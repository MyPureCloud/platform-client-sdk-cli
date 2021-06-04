package conversations_chats

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_communications"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_chats_messages"
)

func init() {
	conversations_chatsCmd.AddCommand(conversations_chats_participants.Cmdconversations_chats_participants())
	conversations_chatsCmd.AddCommand(conversations_chats_communications.Cmdconversations_chats_communications())
	conversations_chatsCmd.AddCommand(conversations_chats_messages.Cmdconversations_chats_messages())
	conversations_chatsCmd.Short = utils.GenerateCustomDescription(conversations_chatsCmd.Short, conversations_chats_participants.Description, conversations_chats_communications.Description, conversations_chats_messages.Description, )
	conversations_chatsCmd.Long = conversations_chatsCmd.Short
}
