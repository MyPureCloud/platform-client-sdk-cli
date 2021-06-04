package telephony_providers_edges_diagnostic

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_providers_edges_diagnostic", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/{edgeId}/diagnostic")
	telephony_providers_edges_diagnosticCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_diagnostic"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_diagnosticCmd)
}

func Cmdtelephony_providers_edges_diagnostic() *cobra.Command {
	return telephony_providers_edges_diagnosticCmd
}
