package users_chats

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_chats_me"
)

func init() {
	users_chatsCmd.AddCommand(users_chats_me.Cmdusers_chats_me())
	users_chatsCmd.Short = utils.GenerateCustomDescription(users_chatsCmd.Short, users_chats_me.Description, )
	users_chatsCmd.Long = users_chatsCmd.Short
}
