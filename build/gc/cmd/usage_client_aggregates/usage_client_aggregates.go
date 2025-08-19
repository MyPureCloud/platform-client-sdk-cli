package usage_client_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("usage_client_aggregates", "SWAGGER_OVERRIDE_/api/v2/usage/client/{clientId}/aggregates")
	usage_client_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage_client_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usage_client_aggregatesCmd)
}

func Cmdusage_client_aggregates() *cobra.Command {
	return usage_client_aggregatesCmd
}
