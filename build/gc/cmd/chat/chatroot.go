package chat

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chat_settings"
)

func init() {
	chatCmd.AddCommand(chat_settings.Cmdchat_settings())
	chatCmd.Short = utils.GenerateCustomDescription(chatCmd.Short, chat_settings.Description, )
	chatCmd.Long = chatCmd.Short
}
