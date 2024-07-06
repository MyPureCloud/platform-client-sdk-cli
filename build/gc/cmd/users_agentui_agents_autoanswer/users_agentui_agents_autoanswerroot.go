package users_agentui_agents_autoanswer

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_agentui_agents_autoanswer_settings"
)

func init() {
	users_agentui_agents_autoanswerCmd.AddCommand(users_agentui_agents_autoanswer_settings.Cmdusers_agentui_agents_autoanswer_settings())
	users_agentui_agents_autoanswerCmd.Short = utils.GenerateCustomDescription(users_agentui_agents_autoanswerCmd.Short, users_agentui_agents_autoanswer_settings.Description, )
	users_agentui_agents_autoanswerCmd.Long = users_agentui_agents_autoanswerCmd.Short
}
