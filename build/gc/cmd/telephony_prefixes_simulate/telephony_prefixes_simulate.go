package telephony_prefixes_simulate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_prefixes_simulate", "SWAGGER_OVERRIDE_/api/v2/telephony/prefixes/simulate")
	telephony_prefixes_simulateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_prefixes_simulate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_prefixes_simulateCmd)
}

func Cmdtelephony_prefixes_simulate() *cobra.Command {
	return telephony_prefixes_simulateCmd
}
