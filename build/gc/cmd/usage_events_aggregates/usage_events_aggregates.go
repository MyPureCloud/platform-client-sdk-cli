package usage_events_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("usage_events_aggregates", "SWAGGER_OVERRIDE_/api/v2/usage/events/aggregates")
	usage_events_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage_events_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usage_events_aggregatesCmd)
}

func Cmdusage_events_aggregates() *cobra.Command {
	return usage_events_aggregatesCmd
}
