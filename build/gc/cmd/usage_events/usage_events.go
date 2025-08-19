package usage_events

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("usage_events", "SWAGGER_OVERRIDE_/api/v2/usage/events")
	usage_eventsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("usage_events"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(usage_eventsCmd)
}

func Cmdusage_events() *cobra.Command {
	return usage_eventsCmd
}
