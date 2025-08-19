package conversations_socials

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_socials_recordingstate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_socials_participants"
)

func init() {
	conversations_socialsCmd.AddCommand(conversations_socials_recordingstate.Cmdconversations_socials_recordingstate())
	conversations_socialsCmd.AddCommand(conversations_socials_participants.Cmdconversations_socials_participants())
	conversations_socialsCmd.Short = utils.GenerateCustomDescription(conversations_socialsCmd.Short, conversations_socials_recordingstate.Description, conversations_socials_participants.Description, )
	conversations_socialsCmd.Long = conversations_socialsCmd.Short
}
