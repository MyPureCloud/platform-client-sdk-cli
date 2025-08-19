package integrations_botconnectors_outgoing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_botconnectors_outgoing", "SWAGGER_OVERRIDE_/api/v2/integrations/botconnectors/outgoing")
	integrations_botconnectors_outgoingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_botconnectors_outgoing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_botconnectors_outgoingCmd)
}

func Cmdintegrations_botconnectors_outgoing() *cobra.Command {
	return integrations_botconnectors_outgoingCmd
}
