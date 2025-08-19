package routing_users_directroutingbackup

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_users_directroutingbackup", "SWAGGER_OVERRIDE_/api/v2/routing/users/{userId}/directroutingbackup")
	routing_users_directroutingbackupCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_users_directroutingbackup"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_users_directroutingbackupCmd)
}

func Cmdrouting_users_directroutingbackup() *cobra.Command {
	return routing_users_directroutingbackupCmd
}
