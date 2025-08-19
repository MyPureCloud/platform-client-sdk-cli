package routing_directroutingbackup

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_directroutingbackup", "SWAGGER_OVERRIDE_/api/v2/routing/directroutingbackup")
	routing_directroutingbackupCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_directroutingbackup"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_directroutingbackupCmd)
}

func Cmdrouting_directroutingbackup() *cobra.Command {
	return routing_directroutingbackupCmd
}
