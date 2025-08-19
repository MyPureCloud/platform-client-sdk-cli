package integrations_actions_functions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_actions_functions", "SWAGGER_OVERRIDE_/api/v2/integrations/actions/functions")
	integrations_actions_functionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_actions_functions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_actions_functionsCmd)
}

func Cmdintegrations_actions_functions() *cobra.Command {
	return integrations_actions_functionsCmd
}
