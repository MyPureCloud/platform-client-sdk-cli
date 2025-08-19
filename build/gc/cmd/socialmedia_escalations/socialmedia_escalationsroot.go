package socialmedia_escalations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_escalations_messages"
)

func init() {
	socialmedia_escalationsCmd.AddCommand(socialmedia_escalations_messages.Cmdsocialmedia_escalations_messages())
	socialmedia_escalationsCmd.Short = utils.GenerateCustomDescription(socialmedia_escalationsCmd.Short, socialmedia_escalations_messages.Description, )
	socialmedia_escalationsCmd.Long = socialmedia_escalationsCmd.Short
}
