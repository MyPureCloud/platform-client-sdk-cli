package integrations_speech

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_speech", "SWAGGER_OVERRIDE_/api/v2/integrations/speech")
	integrations_speechCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_speech"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_speechCmd)
}

func Cmdintegrations_speech() *cobra.Command {
	return integrations_speechCmd
}
