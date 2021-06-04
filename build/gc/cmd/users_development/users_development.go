package users_development

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_development", "SWAGGER_OVERRIDE_/api/v2/users/development")
	users_developmentCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_development"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_developmentCmd)
}

func Cmdusers_development() *cobra.Command {
	return users_developmentCmd
}
