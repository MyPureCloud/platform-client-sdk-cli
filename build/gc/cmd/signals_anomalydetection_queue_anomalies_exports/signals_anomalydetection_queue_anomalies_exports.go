package signals_anomalydetection_queue_anomalies_exports

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("signals_anomalydetection_queue_anomalies_exports", "SWAGGER_OVERRIDE_/api/v2/signals/anomalydetection/queue/anomalies/exports")
	signals_anomalydetection_queue_anomalies_exportsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("signals_anomalydetection_queue_anomalies_exports"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(signals_anomalydetection_queue_anomalies_exportsCmd)
}

func Cmdsignals_anomalydetection_queue_anomalies_exports() *cobra.Command {
	return signals_anomalydetection_queue_anomalies_exportsCmd
}
