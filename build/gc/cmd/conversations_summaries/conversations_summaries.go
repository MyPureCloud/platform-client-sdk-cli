package conversations_summaries

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_summaries", "SWAGGER_OVERRIDE_/api/v2/conversations/{conversationId}/summaries")
	conversations_summariesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_summaries"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_summariesCmd)
}

func Cmdconversations_summaries() *cobra.Command {
	return conversations_summariesCmd
}
