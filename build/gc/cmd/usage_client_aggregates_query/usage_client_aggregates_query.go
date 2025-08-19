package usage_client_aggregates_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("usage_client_aggregates_query", "SWAGGER_OVERRIDE_/api/v2/usage/client/{clientId}/aggregates/query")
	usage_client_aggregates_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage_client_aggregates_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usage_client_aggregates_queryCmd)
}

func Cmdusage_client_aggregates_query() *cobra.Command {
	return usage_client_aggregates_queryCmd
}
