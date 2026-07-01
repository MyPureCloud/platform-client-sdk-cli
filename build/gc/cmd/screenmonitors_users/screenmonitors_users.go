package screenmonitors_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("screenmonitors_users", "SWAGGER_OVERRIDE_/api/v2/screenmonitors/users")
	screenmonitors_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("screenmonitors_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(screenmonitors_usersCmd)
}

func Cmdscreenmonitors_users() *cobra.Command {
	return screenmonitors_usersCmd
}
