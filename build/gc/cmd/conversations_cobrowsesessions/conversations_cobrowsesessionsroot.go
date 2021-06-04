package conversations_cobrowsesessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions_participants"
)

func init() {
	conversations_cobrowsesessionsCmd.AddCommand(conversations_cobrowsesessions_participants.Cmdconversations_cobrowsesessions_participants())
	conversations_cobrowsesessionsCmd.Short = utils.GenerateCustomDescription(conversations_cobrowsesessionsCmd.Short, conversations_cobrowsesessions_participants.Description, )
	conversations_cobrowsesessionsCmd.Long = conversations_cobrowsesessionsCmd.Short
}
