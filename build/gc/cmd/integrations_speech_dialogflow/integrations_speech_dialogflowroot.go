package integrations_speech_dialogflow

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_dialogflow_agents"
)

func init() {
	integrations_speech_dialogflowCmd.AddCommand(integrations_speech_dialogflow_agents.Cmdintegrations_speech_dialogflow_agents())
	integrations_speech_dialogflowCmd.Short = utils.GenerateCustomDescription(integrations_speech_dialogflowCmd.Short, integrations_speech_dialogflow_agents.Description, )
	integrations_speech_dialogflowCmd.Long = integrations_speech_dialogflowCmd.Short
}
