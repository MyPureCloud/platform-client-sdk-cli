package integrations_speech_nuance_bots_launch

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_nuance_bots_launch", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/nuance/{nuanceIntegrationId}/bots/launch")
	integrations_speech_nuance_bots_launchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_nuance_bots_launch"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_nuance_bots_launchCmd)
}

func Cmdintegrations_speech_nuance_bots_launch() *cobra.Command {
	return integrations_speech_nuance_bots_launchCmd
}
