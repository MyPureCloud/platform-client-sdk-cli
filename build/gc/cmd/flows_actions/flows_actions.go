package flows_actions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_actions", "SWAGGER_OVERRIDE_/api/v2/flows/actions")
	flows_actionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_actions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_actionsCmd)
}

func Cmdflows_actions() *cobra.Command {
	return flows_actionsCmd
}
