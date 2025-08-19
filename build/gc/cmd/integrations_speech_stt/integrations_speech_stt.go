package integrations_speech_stt

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_stt", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/stt")
	integrations_speech_sttCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_stt"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_sttCmd)
}

func Cmdintegrations_speech_stt() *cobra.Command {
	return integrations_speech_sttCmd
}
