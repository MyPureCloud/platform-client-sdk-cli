package events_routing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/events_routing_customkpiattributions"
)

func init() {
	events_routingCmd.AddCommand(events_routing_customkpiattributions.Cmdevents_routing_customkpiattributions())
	events_routingCmd.Short = utils.GenerateCustomDescription(events_routingCmd.Short, events_routing_customkpiattributions.Description, )
	events_routingCmd.Long = events_routingCmd.Short
}
