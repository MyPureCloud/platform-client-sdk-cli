package conversations_calls_participants_snippet

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_calls_participants_snippet", "SWAGGER_OVERRIDE_/api/v2/conversations/calls/{conversationId}/participants/{participantId}/snippet")
	conversations_calls_participants_snippetCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_calls_participants_snippet"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_calls_participants_snippetCmd)
}

func Cmdconversations_calls_participants_snippet() *cobra.Command {
	return conversations_calls_participants_snippetCmd
}
