package presence_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("presence_users", "SWAGGER_OVERRIDE_/api/v2/presence/users")
	presence_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("presence_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(presence_usersCmd)
}

func Cmdpresence_users() *cobra.Command {
	return presence_usersCmd
}
