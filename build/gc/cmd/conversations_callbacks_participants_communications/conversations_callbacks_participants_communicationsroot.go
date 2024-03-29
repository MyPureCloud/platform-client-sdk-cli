package conversations_callbacks_participants_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_callbacks_participants_communications_wrapup"
)

func init() {
	conversations_callbacks_participants_communicationsCmd.AddCommand(conversations_callbacks_participants_communications_wrapup.Cmdconversations_callbacks_participants_communications_wrapup())
	conversations_callbacks_participants_communicationsCmd.Short = utils.GenerateCustomDescription(conversations_callbacks_participants_communicationsCmd.Short, conversations_callbacks_participants_communications_wrapup.Description, )
	conversations_callbacks_participants_communicationsCmd.Long = conversations_callbacks_participants_communicationsCmd.Short
}
