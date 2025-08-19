package users_agentui_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_agentui_agents_autoanswer"
)

func init() {
	users_agentui_agentsCmd.AddCommand(users_agentui_agents_autoanswer.Cmdusers_agentui_agents_autoanswer())
	users_agentui_agentsCmd.Short = utils.GenerateCustomDescription(users_agentui_agentsCmd.Short, users_agentui_agents_autoanswer.Description, )
	users_agentui_agentsCmd.Long = users_agentui_agentsCmd.Short
}
