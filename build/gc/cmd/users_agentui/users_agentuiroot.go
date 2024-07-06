package users_agentui

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_agentui_agents"
)

func init() {
	users_agentuiCmd.AddCommand(users_agentui_agents.Cmdusers_agentui_agents())
	users_agentuiCmd.Short = utils.GenerateCustomDescription(users_agentuiCmd.Short, users_agentui_agents.Description, )
	users_agentuiCmd.Long = users_agentuiCmd.Short
}
