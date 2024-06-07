package analytics_flows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows_observations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows_activity"
)

func init() {
	analytics_flowsCmd.AddCommand(analytics_flows_aggregates.Cmdanalytics_flows_aggregates())
	analytics_flowsCmd.AddCommand(analytics_flows_observations.Cmdanalytics_flows_observations())
	analytics_flowsCmd.AddCommand(analytics_flows_activity.Cmdanalytics_flows_activity())
	analytics_flowsCmd.Short = utils.GenerateCustomDescription(analytics_flowsCmd.Short, analytics_flows_aggregates.Description, analytics_flows_observations.Description, analytics_flows_activity.Description, )
	analytics_flowsCmd.Long = analytics_flowsCmd.Short
}
