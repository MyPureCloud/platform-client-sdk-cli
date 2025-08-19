package analytics_summaries_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_summaries_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/summaries/aggregates")
	analytics_summaries_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_summaries_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_summaries_aggregatesCmd)
}

func Cmdanalytics_summaries_aggregates() *cobra.Command {
	return analytics_summaries_aggregatesCmd
}
