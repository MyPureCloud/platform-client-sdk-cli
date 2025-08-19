package webmessaging

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webmessaging_messages"
)

func init() {
	webmessagingCmd.AddCommand(webmessaging_messages.Cmdwebmessaging_messages())
	webmessagingCmd.Short = utils.GenerateCustomDescription(webmessagingCmd.Short, webmessaging_messages.Description, )
	webmessagingCmd.Long = webmessagingCmd.Short
}
