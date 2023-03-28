package chats

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_settings"
)

func init() {
	chatsCmd.AddCommand(chats_settings.Cmdchats_settings())
	chatsCmd.Short = utils.GenerateCustomDescription(chatsCmd.Short, chats_settings.Description, )
	chatsCmd.Long = chatsCmd.Short
}
