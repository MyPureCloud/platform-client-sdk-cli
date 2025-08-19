package chats_users_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_users_messages_pins"
)

func init() {
	chats_users_messagesCmd.AddCommand(chats_users_messages_pins.Cmdchats_users_messages_pins())
	chats_users_messagesCmd.Short = utils.GenerateCustomDescription(chats_users_messagesCmd.Short, chats_users_messages_pins.Description, )
	chats_users_messagesCmd.Long = chats_users_messagesCmd.Short
}
