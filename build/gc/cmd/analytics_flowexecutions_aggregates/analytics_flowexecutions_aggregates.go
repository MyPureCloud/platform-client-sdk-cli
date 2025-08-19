package analytics_flowexecutions_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_flowexecutions_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/flowexecutions/aggregates")
	analytics_flowexecutions_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_flowexecutions_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_flowexecutions_aggregatesCmd)
}

func Cmdanalytics_flowexecutions_aggregates() *cobra.Command {
	return analytics_flowexecutions_aggregatesCmd
}
