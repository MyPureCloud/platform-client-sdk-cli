package analytics_transcripts_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_transcripts_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/transcripts/aggregates")
	analytics_transcripts_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_transcripts_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_transcripts_aggregatesCmd)
}

func Cmdanalytics_transcripts_aggregates() *cobra.Command {
	return analytics_transcripts_aggregatesCmd
}
