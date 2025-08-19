package greetings_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("greetings_users", "SWAGGER_OVERRIDE_/api/v2/greetings/{greetingId}/users")
	greetings_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("greetings_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(greetings_usersCmd)
}

func Cmdgreetings_users() *cobra.Command {
	return greetings_usersCmd
}
