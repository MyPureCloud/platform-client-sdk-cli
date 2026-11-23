package analytics_kpiattributions_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_kpiattributions_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/kpiattributions/aggregates")
	analytics_kpiattributions_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_kpiattributions_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_kpiattributions_aggregatesCmd)
}

func Cmdanalytics_kpiattributions_aggregates() *cobra.Command {
	return analytics_kpiattributions_aggregatesCmd
}
