package events

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("events", "SWAGGER_OVERRIDE_/api/v2/events")
	eventsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("events"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(eventsCmd)
}

func Cmdevents() *cobra.Command {
	return eventsCmd
}
