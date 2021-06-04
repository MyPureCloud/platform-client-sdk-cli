package integrations_speech_tts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_tts_engines"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_tts_settings"
)

func init() {
	integrations_speech_ttsCmd.AddCommand(integrations_speech_tts_engines.Cmdintegrations_speech_tts_engines())
	integrations_speech_ttsCmd.AddCommand(integrations_speech_tts_settings.Cmdintegrations_speech_tts_settings())
	integrations_speech_ttsCmd.Short = utils.GenerateCustomDescription(integrations_speech_ttsCmd.Short, integrations_speech_tts_engines.Description, integrations_speech_tts_settings.Description, )
	integrations_speech_ttsCmd.Long = integrations_speech_ttsCmd.Short
}
