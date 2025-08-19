package integrations_speech_lex_bot

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_lex_bot", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/lex/bot")
	integrations_speech_lex_botCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_lex_bot"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_lex_botCmd)
}

func Cmdintegrations_speech_lex_bot() *cobra.Command {
	return integrations_speech_lex_botCmd
}
