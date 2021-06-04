package analytics_queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_queues", "SWAGGER_OVERRIDE_/api/v2/analytics/queues")
	analytics_queuesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_queues"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_queuesCmd)
}

func Cmdanalytics_queues() *cobra.Command {
	return analytics_queuesCmd
}
