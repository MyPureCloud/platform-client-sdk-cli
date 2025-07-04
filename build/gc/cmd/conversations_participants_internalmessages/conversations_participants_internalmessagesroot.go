package conversations_participants_internalmessages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_participants_internalmessages_users"
)

func init() {
	conversations_participants_internalmessagesCmd.AddCommand(conversations_participants_internalmessages_users.Cmdconversations_participants_internalmessages_users())
	conversations_participants_internalmessagesCmd.Short = utils.GenerateCustomDescription(conversations_participants_internalmessagesCmd.Short, conversations_participants_internalmessages_users.Description, )
	conversations_participants_internalmessagesCmd.Long = conversations_participants_internalmessagesCmd.Short
}
