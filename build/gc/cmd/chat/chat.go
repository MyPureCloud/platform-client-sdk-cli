package chat

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("chat", "SWAGGER_OVERRIDE_/api/v2/chat")
	chatCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("chat"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(chatCmd)
}

func Cmdchat() *cobra.Command {
	return chatCmd
}
