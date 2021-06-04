package analytics_flows_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows_aggregates_query"
)

func init() {
	analytics_flows_aggregatesCmd.AddCommand(analytics_flows_aggregates_query.Cmdanalytics_flows_aggregates_query())
	analytics_flows_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_flows_aggregatesCmd.Short, analytics_flows_aggregates_query.Description, )
	analytics_flows_aggregatesCmd.Long = analytics_flows_aggregatesCmd.Short
}
