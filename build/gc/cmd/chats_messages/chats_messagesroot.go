package chats_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_messages_reactions"
)

func init() {
	chats_messagesCmd.AddCommand(chats_messages_reactions.Cmdchats_messages_reactions())
	chats_messagesCmd.Short = utils.GenerateCustomDescription(chats_messagesCmd.Short, chats_messages_reactions.Description, )
	chats_messagesCmd.Long = chats_messagesCmd.Short
}
