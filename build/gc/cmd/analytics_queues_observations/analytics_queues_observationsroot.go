package analytics_queues_observations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_queues_observations_query"
)

func init() {
	analytics_queues_observationsCmd.AddCommand(analytics_queues_observations_query.Cmdanalytics_queues_observations_query())
	analytics_queues_observationsCmd.Short = utils.GenerateCustomDescription(analytics_queues_observationsCmd.Short, analytics_queues_observations_query.Description, )
	analytics_queues_observationsCmd.Long = analytics_queues_observationsCmd.Short
}
