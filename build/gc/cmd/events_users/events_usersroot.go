package events_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/events_users_routingstatus"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/events_users_presence"
)

func init() {
	events_usersCmd.AddCommand(events_users_routingstatus.Cmdevents_users_routingstatus())
	events_usersCmd.AddCommand(events_users_presence.Cmdevents_users_presence())
	events_usersCmd.Short = utils.GenerateCustomDescription(events_usersCmd.Short, events_users_routingstatus.Description, events_users_presence.Description, )
	events_usersCmd.Long = events_usersCmd.Short
}
