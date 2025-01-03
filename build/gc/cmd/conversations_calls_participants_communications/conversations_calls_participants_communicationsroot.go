package conversations_calls_participants_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_communications_wrapup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_calls_participants_communications_uuidata"
)

func init() {
	conversations_calls_participants_communicationsCmd.AddCommand(conversations_calls_participants_communications_wrapup.Cmdconversations_calls_participants_communications_wrapup())
	conversations_calls_participants_communicationsCmd.AddCommand(conversations_calls_participants_communications_uuidata.Cmdconversations_calls_participants_communications_uuidata())
	conversations_calls_participants_communicationsCmd.Short = utils.GenerateCustomDescription(conversations_calls_participants_communicationsCmd.Short, conversations_calls_participants_communications_wrapup.Description, conversations_calls_participants_communications_uuidata.Description, )
	conversations_calls_participants_communicationsCmd.Long = conversations_calls_participants_communicationsCmd.Short
}
