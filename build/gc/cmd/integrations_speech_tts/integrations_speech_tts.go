package integrations_speech_tts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_tts", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/tts")
	integrations_speech_ttsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_tts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_ttsCmd)
}

func Cmdintegrations_speech_tts() *cobra.Command {
	return integrations_speech_ttsCmd
}
