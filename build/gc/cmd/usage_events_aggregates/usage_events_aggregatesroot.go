package usage_events_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/usage_events_aggregates_query"
)

func init() {
	usage_events_aggregatesCmd.AddCommand(usage_events_aggregates_query.Cmdusage_events_aggregates_query())
	usage_events_aggregatesCmd.Short = utils.GenerateCustomDescription(usage_events_aggregatesCmd.Short, usage_events_aggregates_query.Description, )
	usage_events_aggregatesCmd.Long = usage_events_aggregatesCmd.Short
}
