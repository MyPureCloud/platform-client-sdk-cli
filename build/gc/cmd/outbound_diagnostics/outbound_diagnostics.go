package outbound_diagnostics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_diagnostics", "SWAGGER_OVERRIDE_/api/v2/outbound/diagnostics")
	outbound_diagnosticsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_diagnostics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_diagnosticsCmd)
}

func Cmdoutbound_diagnostics() *cobra.Command {
	return outbound_diagnosticsCmd
}
