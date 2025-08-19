package usage_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("usage_aggregates", "SWAGGER_OVERRIDE_/api/v2/usage/aggregates")
	usage_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usage_aggregatesCmd)
}

func Cmdusage_aggregates() *cobra.Command {
	return usage_aggregatesCmd
}
