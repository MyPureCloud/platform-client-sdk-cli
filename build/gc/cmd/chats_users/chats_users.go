package chats_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("chats_users", "SWAGGER_OVERRIDE_/api/v2/chats/users")
	chats_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("chats_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(chats_usersCmd)
}

func Cmdchats_users() *cobra.Command {
	return chats_usersCmd
}
