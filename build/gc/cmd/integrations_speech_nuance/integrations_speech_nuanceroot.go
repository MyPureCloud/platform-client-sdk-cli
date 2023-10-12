package integrations_speech_nuance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_nuance_bots"
)

func init() {
	integrations_speech_nuanceCmd.AddCommand(integrations_speech_nuance_bots.Cmdintegrations_speech_nuance_bots())
	integrations_speech_nuanceCmd.Short = utils.GenerateCustomDescription(integrations_speech_nuanceCmd.Short, integrations_speech_nuance_bots.Description, )
	integrations_speech_nuanceCmd.Long = integrations_speech_nuanceCmd.Short
}
