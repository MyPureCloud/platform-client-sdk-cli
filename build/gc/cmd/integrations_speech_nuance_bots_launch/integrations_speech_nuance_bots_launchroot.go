package integrations_speech_nuance_bots_launch

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_nuance_bots_launch_validate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_nuance_bots_launch_settings"
)

func init() {
	integrations_speech_nuance_bots_launchCmd.AddCommand(integrations_speech_nuance_bots_launch_validate.Cmdintegrations_speech_nuance_bots_launch_validate())
	integrations_speech_nuance_bots_launchCmd.AddCommand(integrations_speech_nuance_bots_launch_settings.Cmdintegrations_speech_nuance_bots_launch_settings())
	integrations_speech_nuance_bots_launchCmd.Short = utils.GenerateCustomDescription(integrations_speech_nuance_bots_launchCmd.Short, integrations_speech_nuance_bots_launch_validate.Description, integrations_speech_nuance_bots_launch_settings.Description, )
	integrations_speech_nuance_bots_launchCmd.Long = integrations_speech_nuance_bots_launchCmd.Short
}
