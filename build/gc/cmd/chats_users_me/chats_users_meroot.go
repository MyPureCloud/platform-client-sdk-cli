package chats_users_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_users_me_settings"
)

func init() {
	chats_users_meCmd.AddCommand(chats_users_me_settings.Cmdchats_users_me_settings())
	chats_users_meCmd.Short = utils.GenerateCustomDescription(chats_users_meCmd.Short, chats_users_me_settings.Description, )
	chats_users_meCmd.Long = chats_users_meCmd.Short
}
