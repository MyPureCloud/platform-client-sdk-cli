package analytics_routing_activity

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_routing_activity_query"
)

func init() {
	analytics_routing_activityCmd.AddCommand(analytics_routing_activity_query.Cmdanalytics_routing_activity_query())
	analytics_routing_activityCmd.Short = utils.GenerateCustomDescription(analytics_routing_activityCmd.Short, analytics_routing_activity_query.Description, )
	analytics_routing_activityCmd.Long = analytics_routing_activityCmd.Short
}
