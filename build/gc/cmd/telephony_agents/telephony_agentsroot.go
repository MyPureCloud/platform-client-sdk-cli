package telephony_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_agents_greetings"
)

func init() {
	telephony_agentsCmd.AddCommand(telephony_agents_greetings.Cmdtelephony_agents_greetings())
	telephony_agentsCmd.Short = utils.GenerateCustomDescription(telephony_agentsCmd.Short, telephony_agents_greetings.Description, )
	telephony_agentsCmd.Long = telephony_agentsCmd.Short
}
