package events_routing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("events_routing", "SWAGGER_OVERRIDE_/api/v2/events/routing")
	events_routingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("events_routing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(events_routingCmd)
}

func Cmdevents_routing() *cobra.Command {
	return events_routingCmd
}
