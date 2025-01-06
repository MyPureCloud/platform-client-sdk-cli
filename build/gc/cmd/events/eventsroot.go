package events

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/events_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/events_conversations"
)

func init() {
	eventsCmd.AddCommand(events_users.Cmdevents_users())
	eventsCmd.AddCommand(events_conversations.Cmdevents_conversations())
	eventsCmd.Short = utils.GenerateCustomDescription(eventsCmd.Short, events_users.Description, events_conversations.Description, )
	eventsCmd.Long = eventsCmd.Short
}
