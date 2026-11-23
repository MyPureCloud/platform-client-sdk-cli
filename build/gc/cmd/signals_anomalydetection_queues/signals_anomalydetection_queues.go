package signals_anomalydetection_queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("signals_anomalydetection_queues", "SWAGGER_OVERRIDE_/api/v2/signals/anomalydetection/queues")
	signals_anomalydetection_queuesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("signals_anomalydetection_queues"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(signals_anomalydetection_queuesCmd)
}

func Cmdsignals_anomalydetection_queues() *cobra.Command {
	return signals_anomalydetection_queuesCmd
}
