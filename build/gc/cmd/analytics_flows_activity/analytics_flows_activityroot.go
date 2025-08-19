package analytics_flows_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows_activity_query"
)

func init() {
	analytics_flows_activityCmd.AddCommand(analytics_flows_activity_query.Cmdanalytics_flows_activity_query())
	analytics_flows_activityCmd.Short = utils.GenerateCustomDescription(analytics_flows_activityCmd.Short, analytics_flows_activity_query.Description, )
	analytics_flows_activityCmd.Long = analytics_flows_activityCmd.Short
}
