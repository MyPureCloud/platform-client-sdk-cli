package signals_anomalydetection

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("signals_anomalydetection", "SWAGGER_OVERRIDE_/api/v2/signals/anomalydetection")
	signals_anomalydetectionCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("signals_anomalydetection"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(signals_anomalydetectionCmd)
}

func Cmdsignals_anomalydetection() *cobra.Command {
	return signals_anomalydetectionCmd
}
