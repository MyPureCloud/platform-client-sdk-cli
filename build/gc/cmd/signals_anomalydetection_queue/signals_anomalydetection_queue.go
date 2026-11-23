package signals_anomalydetection_queue

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("signals_anomalydetection_queue", "SWAGGER_OVERRIDE_/api/v2/signals/anomalydetection/queue")
	signals_anomalydetection_queueCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("signals_anomalydetection_queue"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(signals_anomalydetection_queueCmd)
}

func Cmdsignals_anomalydetection_queue() *cobra.Command {
	return signals_anomalydetection_queueCmd
}
