package learning_modules_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("learning_modules_users", "SWAGGER_OVERRIDE_/api/v2/learning/modules/{moduleId}/users")
	learning_modules_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("learning_modules_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(learning_modules_usersCmd)
}

func Cmdlearning_modules_users() *cobra.Command {
	return learning_modules_usersCmd
}
