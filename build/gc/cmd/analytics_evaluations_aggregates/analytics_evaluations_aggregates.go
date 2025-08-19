package analytics_evaluations_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_evaluations_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/evaluations/aggregates")
	analytics_evaluations_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_evaluations_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_evaluations_aggregatesCmd)
}

func Cmdanalytics_evaluations_aggregates() *cobra.Command {
	return analytics_evaluations_aggregatesCmd
}
