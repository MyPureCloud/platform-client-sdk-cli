package telephony_providers_edges_mediastatistics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_providers_edges_mediastatistics", "SWAGGER_OVERRIDE_/api/v2/telephony/providers/edges/mediastatistics")
	telephony_providers_edges_mediastatisticsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_providers_edges_mediastatistics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_providers_edges_mediastatisticsCmd)
}

func Cmdtelephony_providers_edges_mediastatistics() *cobra.Command {
	return telephony_providers_edges_mediastatisticsCmd
}
