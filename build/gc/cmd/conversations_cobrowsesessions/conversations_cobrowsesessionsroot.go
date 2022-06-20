package conversations_cobrowsesessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_cobrowsesessions_recordingstate"
)

func init() {
	conversations_cobrowsesessionsCmd.AddCommand(conversations_cobrowsesessions_participants.Cmdconversations_cobrowsesessions_participants())
	conversations_cobrowsesessionsCmd.AddCommand(conversations_cobrowsesessions_recordingstate.Cmdconversations_cobrowsesessions_recordingstate())
	conversations_cobrowsesessionsCmd.Short = utils.GenerateCustomDescription(conversations_cobrowsesessionsCmd.Short, conversations_cobrowsesessions_participants.Description, conversations_cobrowsesessions_recordingstate.Description, )
	conversations_cobrowsesessionsCmd.Long = conversations_cobrowsesessionsCmd.Short
}
