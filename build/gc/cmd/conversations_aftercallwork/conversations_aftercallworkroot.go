package conversations_aftercallwork

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_aftercallwork_participants"
)

func init() {
	conversations_aftercallworkCmd.AddCommand(conversations_aftercallwork_participants.Cmdconversations_aftercallwork_participants())
	conversations_aftercallworkCmd.Short = utils.GenerateCustomDescription(conversations_aftercallworkCmd.Short, conversations_aftercallwork_participants.Description, )
	conversations_aftercallworkCmd.Long = conversations_aftercallworkCmd.Short
}
