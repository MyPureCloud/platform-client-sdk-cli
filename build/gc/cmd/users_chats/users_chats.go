package users_chats

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_chats", "SWAGGER_OVERRIDE_/api/v2/users/chats")
	users_chatsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_chats"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_chatsCmd)
}

func Cmdusers_chats() *cobra.Command {
	return users_chatsCmd
}
