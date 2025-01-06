package analytics_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agents_status"
)

func init() {
	analytics_agentsCmd.AddCommand(analytics_agents_status.Cmdanalytics_agents_status())
	analytics_agentsCmd.Short = utils.GenerateCustomDescription(analytics_agentsCmd.Short, analytics_agents_status.Description, )
	analytics_agentsCmd.Long = analytics_agentsCmd.Short
}
