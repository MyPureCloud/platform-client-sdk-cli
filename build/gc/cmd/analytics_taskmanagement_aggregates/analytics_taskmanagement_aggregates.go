package analytics_taskmanagement_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_taskmanagement_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/taskmanagement/aggregates")
	analytics_taskmanagement_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_taskmanagement_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_taskmanagement_aggregatesCmd)
}

func Cmdanalytics_taskmanagement_aggregates() *cobra.Command {
	return analytics_taskmanagement_aggregatesCmd
}
