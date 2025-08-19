package conversations_messaging

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messaging", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging")
	conversations_messagingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messaging"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messagingCmd)
}

func Cmdconversations_messaging() *cobra.Command {
	return conversations_messagingCmd
}
