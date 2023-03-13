package conversations_socials_participants

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_socials_participants_communications"
)

func init() {
	conversations_socials_participantsCmd.AddCommand(conversations_socials_participants_communications.Cmdconversations_socials_participants_communications())
	conversations_socials_participantsCmd.Short = utils.GenerateCustomDescription(conversations_socials_participantsCmd.Short, conversations_socials_participants_communications.Description, )
	conversations_socials_participantsCmd.Long = conversations_socials_participantsCmd.Short
}
