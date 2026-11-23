package signals_anomalydetection_queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/signals_anomalydetection_queues_analysis"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/signals_anomalydetection_queues_available"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/signals_anomalydetection_queues_monitored"
)

func init() {
	signals_anomalydetection_queuesCmd.AddCommand(signals_anomalydetection_queues_analysis.Cmdsignals_anomalydetection_queues_analysis())
	signals_anomalydetection_queuesCmd.AddCommand(signals_anomalydetection_queues_available.Cmdsignals_anomalydetection_queues_available())
	signals_anomalydetection_queuesCmd.AddCommand(signals_anomalydetection_queues_monitored.Cmdsignals_anomalydetection_queues_monitored())
	signals_anomalydetection_queuesCmd.Short = utils.GenerateCustomDescription(signals_anomalydetection_queuesCmd.Short, signals_anomalydetection_queues_analysis.Description, signals_anomalydetection_queues_available.Description, signals_anomalydetection_queues_monitored.Description, )
	signals_anomalydetection_queuesCmd.Long = signals_anomalydetection_queuesCmd.Short
}
