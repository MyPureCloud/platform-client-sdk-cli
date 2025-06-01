package conversations_messages_communications_socialmedia

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messages_communications_socialmedia", "SWAGGER_OVERRIDE_/api/v2/conversations/messages/{conversationId}/communications/{communicationId}/socialmedia")
	conversations_messages_communications_socialmediaCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messages_communications_socialmedia"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messages_communications_socialmediaCmd)
}

func Cmdconversations_messages_communications_socialmedia() *cobra.Command {
	return conversations_messages_communications_socialmediaCmd
}
