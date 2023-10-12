package integrations_speech_nuance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_nuance", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/nuance")
	integrations_speech_nuanceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_nuance"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_nuanceCmd)
}

func Cmdintegrations_speech_nuance() *cobra.Command {
	return integrations_speech_nuanceCmd
}
