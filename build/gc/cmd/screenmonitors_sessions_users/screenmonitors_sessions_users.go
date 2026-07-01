package screenmonitors_sessions_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("screenmonitors_sessions_users", "SWAGGER_OVERRIDE_/api/v2/screenmonitors/sessions/users")
	screenmonitors_sessions_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("screenmonitors_sessions_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(screenmonitors_sessions_usersCmd)
}

func Cmdscreenmonitors_sessions_users() *cobra.Command {
	return screenmonitors_sessions_usersCmd
}
