package chats

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("chats", "SWAGGER_OVERRIDE_/api/v2/chats")
	chatsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("chats"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(chatsCmd)
}

func Cmdchats() *cobra.Command {
	return chatsCmd
}
