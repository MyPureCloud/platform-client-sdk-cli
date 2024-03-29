package conversations_messages_participants_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messages_participants_communications_wrapup"
)

func init() {
	conversations_messages_participants_communicationsCmd.AddCommand(conversations_messages_participants_communications_wrapup.Cmdconversations_messages_participants_communications_wrapup())
	conversations_messages_participants_communicationsCmd.Short = utils.GenerateCustomDescription(conversations_messages_participants_communicationsCmd.Short, conversations_messages_participants_communications_wrapup.Description, )
	conversations_messages_participants_communicationsCmd.Long = conversations_messages_participants_communicationsCmd.Short
}
