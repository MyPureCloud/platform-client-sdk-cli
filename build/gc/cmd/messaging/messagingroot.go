package messaging

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/messaging_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/messaging_supportedcontent"
)

func init() {
	messagingCmd.AddCommand(messaging_settings.Cmdmessaging_settings())
	messagingCmd.AddCommand(messaging_supportedcontent.Cmdmessaging_supportedcontent())
	messagingCmd.Short = utils.GenerateCustomDescription(messagingCmd.Short, messaging_settings.Description, messaging_supportedcontent.Description, )
	messagingCmd.Long = messagingCmd.Short
}
