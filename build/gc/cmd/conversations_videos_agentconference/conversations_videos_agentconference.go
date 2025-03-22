package conversations_videos_agentconference

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_videos_agentconference", "SWAGGER_OVERRIDE_/api/v2/conversations/videos/{conversationId}/agentconference")
	conversations_videos_agentconferenceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_videos_agentconference"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_videos_agentconferenceCmd)
}

func Cmdconversations_videos_agentconference() *cobra.Command {
	return conversations_videos_agentconferenceCmd
}
