package integrations_speech_nuance_bots

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_nuance_bots_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_nuance_bots_launch"
)

func init() {
	integrations_speech_nuance_botsCmd.AddCommand(integrations_speech_nuance_bots_jobs.Cmdintegrations_speech_nuance_bots_jobs())
	integrations_speech_nuance_botsCmd.AddCommand(integrations_speech_nuance_bots_launch.Cmdintegrations_speech_nuance_bots_launch())
	integrations_speech_nuance_botsCmd.Short = utils.GenerateCustomDescription(integrations_speech_nuance_botsCmd.Short, integrations_speech_nuance_bots_jobs.Description, integrations_speech_nuance_bots_launch.Description, )
	integrations_speech_nuance_botsCmd.Long = integrations_speech_nuance_botsCmd.Short
}
