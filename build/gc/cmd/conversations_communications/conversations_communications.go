package conversations_communications

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_communications", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/communications")
	conversations_communicationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_communications"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_communicationsCmd)
}

func Cmdconversations_communications() *cobra.Command {
	return conversations_communicationsCmd
}
