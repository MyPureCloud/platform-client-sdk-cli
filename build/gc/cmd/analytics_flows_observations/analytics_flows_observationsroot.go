package analytics_flows_observations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows_observations_query"
)

func init() {
	analytics_flows_observationsCmd.AddCommand(analytics_flows_observations_query.Cmdanalytics_flows_observations_query())
	analytics_flows_observationsCmd.Short = utils.GenerateCustomDescription(analytics_flows_observationsCmd.Short, analytics_flows_observations_query.Description, )
	analytics_flows_observationsCmd.Long = analytics_flows_observationsCmd.Short
}
