package quality_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/quality_agents_activity"
)

func init() {
	quality_agentsCmd.AddCommand(quality_agents_activity.Cmdquality_agents_activity())
	quality_agentsCmd.Short = utils.GenerateCustomDescription(quality_agentsCmd.Short, quality_agents_activity.Description, )
	quality_agentsCmd.Long = quality_agentsCmd.Short
}
