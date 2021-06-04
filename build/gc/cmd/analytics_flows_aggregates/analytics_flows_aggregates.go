package analytics_flows_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_flows_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/flows/aggregates")
	analytics_flows_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_flows_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_flows_aggregatesCmd)
}

func Cmdanalytics_flows_aggregates() *cobra.Command {
	return analytics_flows_aggregatesCmd
}
