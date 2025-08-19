package analytics_transcripts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_transcripts", "SWAGGER_OVERRIDE_/api/v2/analytics/transcripts")
	analytics_transcriptsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_transcripts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_transcriptsCmd)
}

func Cmdanalytics_transcripts() *cobra.Command {
	return analytics_transcriptsCmd
}
