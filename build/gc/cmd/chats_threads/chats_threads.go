package chats_threads

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("chats_threads", "SWAGGER_OVERRIDE_/api/v2/chats/threads")
	chats_threadsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("chats_threads"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(chats_threadsCmd)
}

func Cmdchats_threads() *cobra.Command {
	return chats_threadsCmd
}
