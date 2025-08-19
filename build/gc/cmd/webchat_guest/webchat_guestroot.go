package webchat_guest

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webchat_guest_conversations"
)

func init() {
	webchat_guestCmd.AddCommand(webchat_guest_conversations.Cmdwebchat_guest_conversations())
	webchat_guestCmd.Short = utils.GenerateCustomDescription(webchat_guestCmd.Short, webchat_guest_conversations.Description, )
	webchat_guestCmd.Long = webchat_guestCmd.Short
}
