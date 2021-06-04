package analytics_journeys_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_journeys_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/journeys/aggregates")
	analytics_journeys_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_journeys_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_journeys_aggregatesCmd)
}

func Cmdanalytics_journeys_aggregates() *cobra.Command {
	return analytics_journeys_aggregatesCmd
}
