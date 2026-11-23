package signals

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/signals_anomalydetection"
)

func init() {
	signalsCmd.AddCommand(signals_anomalydetection.Cmdsignals_anomalydetection())
	signalsCmd.Short = utils.GenerateCustomDescription(signalsCmd.Short, signals_anomalydetection.Description, )
	signalsCmd.Long = signalsCmd.Short
}
