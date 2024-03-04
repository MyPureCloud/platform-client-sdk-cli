package integrations_speech_dialogflowcx

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_dialogflowcx_agents"
)

func init() {
	integrations_speech_dialogflowcxCmd.AddCommand(integrations_speech_dialogflowcx_agents.Cmdintegrations_speech_dialogflowcx_agents())
	integrations_speech_dialogflowcxCmd.Short = utils.GenerateCustomDescription(integrations_speech_dialogflowcxCmd.Short, integrations_speech_dialogflowcx_agents.Description, )
	integrations_speech_dialogflowcxCmd.Long = integrations_speech_dialogflowcxCmd.Short
}
