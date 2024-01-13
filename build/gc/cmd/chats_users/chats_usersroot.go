package chats_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_users_messages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_users_settings"
)

func init() {
	chats_usersCmd.AddCommand(chats_users_messages.Cmdchats_users_messages())
	chats_usersCmd.AddCommand(chats_users_settings.Cmdchats_users_settings())
	chats_usersCmd.Short = utils.GenerateCustomDescription(chats_usersCmd.Short, chats_users_messages.Description, chats_users_settings.Description, )
	chats_usersCmd.Long = chats_usersCmd.Short
}
