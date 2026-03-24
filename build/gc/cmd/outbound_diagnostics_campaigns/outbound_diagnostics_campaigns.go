package outbound_diagnostics_campaigns

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("outbound_diagnostics_campaigns", "SWAGGER_OVERRIDE_/api/v2/outbound/diagnostics/campaigns")
	outbound_diagnostics_campaignsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("outbound_diagnostics_campaigns"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(outbound_diagnostics_campaignsCmd)
}

func Cmdoutbound_diagnostics_campaigns() *cobra.Command {
	return outbound_diagnostics_campaignsCmd
}
