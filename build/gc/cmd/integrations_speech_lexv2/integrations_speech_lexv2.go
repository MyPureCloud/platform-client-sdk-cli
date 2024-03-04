package integrations_speech_lexv2

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_lexv2", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/lexv2")
	integrations_speech_lexv2Cmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_lexv2"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_lexv2Cmd)
}

func Cmdintegrations_speech_lexv2() *cobra.Command {
	return integrations_speech_lexv2Cmd
}
