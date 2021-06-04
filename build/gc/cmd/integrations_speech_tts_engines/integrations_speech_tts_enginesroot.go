package integrations_speech_tts_engines

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_speech_tts_engines_voices"
)

func init() {
	integrations_speech_tts_enginesCmd.AddCommand(integrations_speech_tts_engines_voices.Cmdintegrations_speech_tts_engines_voices())
	integrations_speech_tts_enginesCmd.Short = utils.GenerateCustomDescription(integrations_speech_tts_enginesCmd.Short, integrations_speech_tts_engines_voices.Description, )
	integrations_speech_tts_enginesCmd.Long = integrations_speech_tts_enginesCmd.Short
}
