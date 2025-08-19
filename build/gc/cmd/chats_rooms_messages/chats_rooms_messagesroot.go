package chats_rooms_messages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_rooms_messages_pins"
)

func init() {
	chats_rooms_messagesCmd.AddCommand(chats_rooms_messages_pins.Cmdchats_rooms_messages_pins())
	chats_rooms_messagesCmd.Short = utils.GenerateCustomDescription(chats_rooms_messagesCmd.Short, chats_rooms_messages_pins.Description, )
	chats_rooms_messagesCmd.Long = chats_rooms_messagesCmd.Short
}
