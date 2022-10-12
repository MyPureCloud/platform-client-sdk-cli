package events

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/events_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/events_users"
)

func init() {
	eventsCmd.AddCommand(events_conversations.Cmdevents_conversations())
	eventsCmd.AddCommand(events_users.Cmdevents_users())
	eventsCmd.Short = utils.GenerateCustomDescription(eventsCmd.Short, events_conversations.Description, events_users.Description, )
	eventsCmd.Long = eventsCmd.Short
}
