package analytics_queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_queues_observations"
)

func init() {
	analytics_queuesCmd.AddCommand(analytics_queues_observations.Cmdanalytics_queues_observations())
	analytics_queuesCmd.Short = utils.GenerateCustomDescription(analytics_queuesCmd.Short, analytics_queues_observations.Description, )
	analytics_queuesCmd.Long = analytics_queuesCmd.Short
}
