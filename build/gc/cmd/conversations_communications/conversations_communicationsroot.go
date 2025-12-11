package conversations_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_communications_agentchecklists"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_communications_internalmessages"
)

func init() {
	conversations_communicationsCmd.AddCommand(conversations_communications_agentchecklists.Cmdconversations_communications_agentchecklists())
	conversations_communicationsCmd.AddCommand(conversations_communications_internalmessages.Cmdconversations_communications_internalmessages())
	conversations_communicationsCmd.Short = utils.GenerateCustomDescription(conversations_communicationsCmd.Short, conversations_communications_agentchecklists.Description, conversations_communications_internalmessages.Description, )
	conversations_communicationsCmd.Long = conversations_communicationsCmd.Short
}
