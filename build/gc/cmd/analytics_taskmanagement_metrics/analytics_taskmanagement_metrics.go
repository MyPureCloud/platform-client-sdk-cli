package analytics_taskmanagement_metrics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_taskmanagement_metrics", "SWAGGER_OVERRIDE_/api/v2/analytics/taskmanagement/metrics")
	analytics_taskmanagement_metricsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_taskmanagement_metrics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_taskmanagement_metricsCmd)
}

func Cmdanalytics_taskmanagement_metrics() *cobra.Command {
	return analytics_taskmanagement_metricsCmd
}
