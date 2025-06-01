package integrations_botconnectors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_botconnectors", "SWAGGER_OVERRIDE_/api/v2/integrations/botconnectors")
	integrations_botconnectorsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_botconnectors"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_botconnectorsCmd)
}

func Cmdintegrations_botconnectors() *cobra.Command {
	return integrations_botconnectorsCmd
}
