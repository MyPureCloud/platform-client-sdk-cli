package analytics_queues_observations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_queues_observations", "SWAGGER_OVERRIDE_/api/v2/analytics/queues/observations")
	analytics_queues_observationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_queues_observations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_queues_observationsCmd)
}

func Cmdanalytics_queues_observations() *cobra.Command {
	return analytics_queues_observationsCmd
}
