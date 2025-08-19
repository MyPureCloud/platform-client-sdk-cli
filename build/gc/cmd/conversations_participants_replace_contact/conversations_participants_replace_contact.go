package conversations_participants_replace_contact

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_participants_replace_contact", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/participants/{participantId}/replace/contact")
	conversations_participants_replace_contactCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_participants_replace_contact"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_participants_replace_contactCmd)
}

func Cmdconversations_participants_replace_contact() *cobra.Command {
	return conversations_participants_replace_contactCmd
}
