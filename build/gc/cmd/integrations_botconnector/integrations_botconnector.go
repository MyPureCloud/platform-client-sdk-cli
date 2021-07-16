package integrations_botconnector

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_botconnector", "SWAGGER_OVERRIDE_/api/v2/integrations/botconnector")
	integrations_botconnectorCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_botconnector"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_botconnectorCmd)
}

func Cmdintegrations_botconnector() *cobra.Command {
	return integrations_botconnectorCmd
}
