package conversations_messages_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messages_communications", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages")
	conversations_messages_communicationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_communications"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_communicationsCmd)
}

func Cmdconversations_messages_communications() *cobra.Command {
	return conversations_messages_communicationsCmd
}
