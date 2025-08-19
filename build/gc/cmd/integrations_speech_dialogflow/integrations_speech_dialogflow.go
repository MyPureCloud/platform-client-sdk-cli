package integrations_speech_dialogflow

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_dialogflow", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/dialogflow")
	integrations_speech_dialogflowCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_dialogflow"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_dialogflowCmd)
}

func Cmdintegrations_speech_dialogflow() *cobra.Command {
	return integrations_speech_dialogflowCmd
}
