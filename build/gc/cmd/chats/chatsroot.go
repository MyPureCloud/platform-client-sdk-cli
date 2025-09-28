package chats

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_rooms"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_threads"
)

func init() {
	chatsCmd.AddCommand(chats_settings.Cmdchats_settings())
	chatsCmd.AddCommand(chats_users.Cmdchats_users())
	chatsCmd.AddCommand(chats_rooms.Cmdchats_rooms())
	chatsCmd.AddCommand(chats_messages.Cmdchats_messages())
	chatsCmd.AddCommand(chats_threads.Cmdchats_threads())
	chatsCmd.Short = utils.GenerateCustomDescription(chatsCmd.Short, chats_settings.Description, chats_users.Description, chats_rooms.Description, chats_messages.Description, chats_threads.Description, )
	chatsCmd.Long = chatsCmd.Short
}
