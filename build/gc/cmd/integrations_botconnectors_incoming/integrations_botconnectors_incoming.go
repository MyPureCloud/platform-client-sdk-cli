package integrations_botconnectors_incoming

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_botconnectors_incoming", "SWAGGER_OVERRIDE_/api/v2/integrations/botconnectors/incoming")
	integrations_botconnectors_incomingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_botconnectors_incoming"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_botconnectors_incomingCmd)
}

func Cmdintegrations_botconnectors_incoming() *cobra.Command {
	return integrations_botconnectors_incomingCmd
}
