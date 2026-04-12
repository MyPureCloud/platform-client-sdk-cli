package analytics_agentutilizations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agentutilizations_aggregates"
)

func init() {
	analytics_agentutilizationsCmd.AddCommand(analytics_agentutilizations_aggregates.Cmdanalytics_agentutilizations_aggregates())
	analytics_agentutilizationsCmd.Short = utils.GenerateCustomDescription(analytics_agentutilizationsCmd.Short, analytics_agentutilizations_aggregates.Description, )
	analytics_agentutilizationsCmd.Long = analytics_agentutilizationsCmd.Short
}
