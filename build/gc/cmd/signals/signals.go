package signals

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("signals", "SWAGGER_OVERRIDE_/api/v2/signals")
	signalsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("signals"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(signalsCmd)
}

func Cmdsignals() *cobra.Command {
	return signalsCmd
}
