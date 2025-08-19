package knowledge_guest

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/knowledge_guest_sessions"
)

func init() {
	knowledge_guestCmd.AddCommand(knowledge_guest_sessions.Cmdknowledge_guest_sessions())
	knowledge_guestCmd.Short = utils.GenerateCustomDescription(knowledge_guestCmd.Short, knowledge_guest_sessions.Description, )
	knowledge_guestCmd.Long = knowledge_guestCmd.Short
}
