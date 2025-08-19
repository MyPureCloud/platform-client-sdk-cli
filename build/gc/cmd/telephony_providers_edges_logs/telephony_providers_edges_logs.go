package telephony_providers_edges_logs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_providers_edges_logs", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/{edgeId}/logs")
	telephony_providers_edges_logsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_logs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_logsCmd)
}

func Cmdtelephony_providers_edges_logs() *cobra.Command {
	return telephony_providers_edges_logsCmd
}
