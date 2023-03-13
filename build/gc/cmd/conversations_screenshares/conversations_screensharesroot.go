package conversations_screenshares

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_screenshares_participants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_screenshares_recordingstate"
)

func init() {
	conversations_screensharesCmd.AddCommand(conversations_screenshares_participants.Cmdconversations_screenshares_participants())
	conversations_screensharesCmd.AddCommand(conversations_screenshares_recordingstate.Cmdconversations_screenshares_recordingstate())
	conversations_screensharesCmd.Short = utils.GenerateCustomDescription(conversations_screensharesCmd.Short, conversations_screenshares_participants.Description, conversations_screenshares_recordingstate.Description, )
	conversations_screensharesCmd.Long = conversations_screensharesCmd.Short
}
