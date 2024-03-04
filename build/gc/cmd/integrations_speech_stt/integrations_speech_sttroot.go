package integrations_speech_stt

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_stt_engines"
)

func init() {
	integrations_speech_sttCmd.AddCommand(integrations_speech_stt_engines.Cmdintegrations_speech_stt_engines())
	integrations_speech_sttCmd.Short = utils.GenerateCustomDescription(integrations_speech_sttCmd.Short, integrations_speech_stt_engines.Description, )
	integrations_speech_sttCmd.Long = integrations_speech_sttCmd.Short
}
