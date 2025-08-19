package conversations_screenshares_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_screenshares_participants_communications"
)

func init() {
	conversations_screenshares_participantsCmd.AddCommand(conversations_screenshares_participants_communications.Cmdconversations_screenshares_participants_communications())
	conversations_screenshares_participantsCmd.Short = utils.GenerateCustomDescription(conversations_screenshares_participantsCmd.Short, conversations_screenshares_participants_communications.Description, )
	conversations_screenshares_participantsCmd.Long = conversations_screenshares_participantsCmd.Short
}
