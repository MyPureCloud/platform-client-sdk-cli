package conversations_callbacks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_callbacks_participants"
)

func init() {
	conversations_callbacksCmd.AddCommand(conversations_callbacks_participants.Cmdconversations_callbacks_participants())
	conversations_callbacksCmd.Short = utils.GenerateCustomDescription(conversations_callbacksCmd.Short, conversations_callbacks_participants.Description, )
	conversations_callbacksCmd.Long = conversations_callbacksCmd.Short
}
