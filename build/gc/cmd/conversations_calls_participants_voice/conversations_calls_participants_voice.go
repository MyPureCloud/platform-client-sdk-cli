package conversations_calls_participants_voice

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_calls_participants_voice", "SWAGGER_OVERRIDE_/api/v2/conversations/calls/{conversationId}/participants/{participantId}/voice")
	conversations_calls_participants_voiceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_calls_participants_voice"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_calls_participants_voiceCmd)
}

func Cmdconversations_calls_participants_voice() *cobra.Command {
	return conversations_calls_participants_voiceCmd
}
