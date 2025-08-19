package routing_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_users", "SWAGGER_OVERRIDE_/api/v2/routing/users")
	routing_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_usersCmd)
}

func Cmdrouting_users() *cobra.Command {
	return routing_usersCmd
}
