package chats_threads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/chats_threads_messages"
)

func init() {
	chats_threadsCmd.AddCommand(chats_threads_messages.Cmdchats_threads_messages())
	chats_threadsCmd.Short = utils.GenerateCustomDescription(chats_threadsCmd.Short, chats_threads_messages.Description, )
	chats_threadsCmd.Long = chats_threadsCmd.Short
}
