package signals_anomalydetection

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/signals_anomalydetection_queues"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/signals_anomalydetection_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/signals_anomalydetection_queue"
)

func init() {
	signals_anomalydetectionCmd.AddCommand(signals_anomalydetection_queues.Cmdsignals_anomalydetection_queues())
	signals_anomalydetectionCmd.AddCommand(signals_anomalydetection_settings.Cmdsignals_anomalydetection_settings())
	signals_anomalydetectionCmd.AddCommand(signals_anomalydetection_queue.Cmdsignals_anomalydetection_queue())
	signals_anomalydetectionCmd.Short = utils.GenerateCustomDescription(signals_anomalydetectionCmd.Short, signals_anomalydetection_queues.Description, signals_anomalydetection_settings.Description, signals_anomalydetection_queue.Description, )
	signals_anomalydetectionCmd.Long = signals_anomalydetectionCmd.Short
}
