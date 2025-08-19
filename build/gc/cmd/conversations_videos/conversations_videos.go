package conversations_videos

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_videos", "SWAGGER_OVERRIDE_/api/v2/conversations/videos")
	conversations_videosCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_videos"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_videosCmd)
}

func Cmdconversations_videos() *cobra.Command {
	return conversations_videosCmd
}
