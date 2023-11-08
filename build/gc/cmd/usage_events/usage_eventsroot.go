package usage_events

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_events_definitions"
)

func init() {
	usage_eventsCmd.AddCommand(usage_events_definitions.Cmdusage_events_definitions())
	usage_eventsCmd.Short = utils.GenerateCustomDescription(usage_eventsCmd.Short, usage_events_definitions.Description, )
	usage_eventsCmd.Long = usage_eventsCmd.Short
}
