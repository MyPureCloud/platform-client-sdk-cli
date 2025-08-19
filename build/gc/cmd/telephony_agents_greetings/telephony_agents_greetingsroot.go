package telephony_agents_greetings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_agents_greetings_me"
)

func init() {
	telephony_agents_greetingsCmd.AddCommand(telephony_agents_greetings_me.Cmdtelephony_agents_greetings_me())
	telephony_agents_greetingsCmd.Short = utils.GenerateCustomDescription(telephony_agents_greetingsCmd.Short, telephony_agents_greetings_me.Description, )
	telephony_agents_greetingsCmd.Long = telephony_agents_greetingsCmd.Short
}
