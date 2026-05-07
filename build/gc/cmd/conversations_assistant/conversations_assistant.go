package conversations_assistant

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_assistant", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/assistant")
	conversations_assistantCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_assistant"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_assistantCmd)
}

func Cmdconversations_assistant() *cobra.Command {
	return conversations_assistantCmd
}
