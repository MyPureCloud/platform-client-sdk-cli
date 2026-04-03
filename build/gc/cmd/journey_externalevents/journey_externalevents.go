package journey_externalevents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_externalevents", "SWAGGER_OVERRIDE_/api/v2/journey/externalevents")
	journey_externaleventsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_externalevents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_externaleventsCmd)
}

func Cmdjourney_externalevents() *cobra.Command {
	return journey_externaleventsCmd
}
