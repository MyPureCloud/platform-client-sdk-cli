package usage_events

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_events_definitions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_events_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_events_aggregates"
)

func init() {
	usage_eventsCmd.AddCommand(usage_events_definitions.Cmdusage_events_definitions())
	usage_eventsCmd.AddCommand(usage_events_query.Cmdusage_events_query())
	usage_eventsCmd.AddCommand(usage_events_aggregates.Cmdusage_events_aggregates())
	usage_eventsCmd.Short = utils.GenerateCustomDescription(usage_eventsCmd.Short, usage_events_definitions.Description, usage_events_query.Description, usage_events_aggregates.Description, )
	usage_eventsCmd.Long = usage_eventsCmd.Short
}
