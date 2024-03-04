package integrations_speech_dialogflowcx

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech_dialogflowcx", "SWAGGER_OVERRIDE_/api/v2/integrations/speech/dialogflowcx")
	integrations_speech_dialogflowcxCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech_dialogflowcx"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speech_dialogflowcxCmd)
}

func Cmdintegrations_speech_dialogflowcx() *cobra.Command {
	return integrations_speech_dialogflowcxCmd
}
