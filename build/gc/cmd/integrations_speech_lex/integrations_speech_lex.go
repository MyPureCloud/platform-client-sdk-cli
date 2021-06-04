package integrations_speech_lex

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_lex", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/lex")
	integrations_speech_lexCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_lex"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_lexCmd)
}

func Cmdintegrations_speech_lex() *cobra.Command {
	return integrations_speech_lexCmd
}
