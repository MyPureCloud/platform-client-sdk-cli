package conversations_calls_participants_consult_contact

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_calls_participants_consult_contact", "SWAGGER_OVERRIDE_/api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult/contact")
	conversations_calls_participants_consult_contactCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_calls_participants_consult_contact"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_calls_participants_consult_contactCmd)
}

func Cmdconversations_calls_participants_consult_contact() *cobra.Command {
	return conversations_calls_participants_consult_contactCmd
}
