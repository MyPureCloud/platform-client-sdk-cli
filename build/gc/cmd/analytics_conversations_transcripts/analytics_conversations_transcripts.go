package analytics_conversations_transcripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_conversations_transcripts", "SWAGGER_OVERRIDE_/api/v2/analytics/conversations/transcripts")
	analytics_conversations_transcriptsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_conversations_transcripts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_conversations_transcriptsCmd)
}

func Cmdanalytics_conversations_transcripts() *cobra.Command {
	return analytics_conversations_transcriptsCmd
}
