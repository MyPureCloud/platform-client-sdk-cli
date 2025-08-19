package conversations_participants_internalmessages

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_participants_internalmessages", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/internalmessages")
	conversations_participants_internalmessagesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_participants_internalmessages"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_participants_internalmessagesCmd)
}

func Cmdconversations_participants_internalmessages() *cobra.Command {
	return conversations_participants_internalmessagesCmd
}
