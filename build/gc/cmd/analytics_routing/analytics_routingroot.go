package analytics_routing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_routing_activity"
)

func init() {
	analytics_routingCmd.AddCommand(analytics_routing_activity.Cmdanalytics_routing_activity())
	analytics_routingCmd.Short = utils.GenerateCustomDescription(analytics_routingCmd.Short, analytics_routing_activity.Description, )
	analytics_routingCmd.Long = analytics_routingCmd.Short
}
