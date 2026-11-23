package signals_anomalydetection_queue

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/signals_anomalydetection_queue_anomalies"
)

func init() {
	signals_anomalydetection_queueCmd.AddCommand(signals_anomalydetection_queue_anomalies.Cmdsignals_anomalydetection_queue_anomalies())
	signals_anomalydetection_queueCmd.Short = utils.GenerateCustomDescription(signals_anomalydetection_queueCmd.Short, signals_anomalydetection_queue_anomalies.Description, )
	signals_anomalydetection_queueCmd.Long = signals_anomalydetection_queueCmd.Short
}
