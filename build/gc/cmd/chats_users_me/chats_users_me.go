package chats_users_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("chats_users_me", "SWAGGER_OVERRIDE_/api/v2/chats/users/me")
	chats_users_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("chats_users_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(chats_users_meCmd)
}

func Cmdchats_users_me() *cobra.Command {
	return chats_users_meCmd
}
