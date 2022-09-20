package journey_actionmaps_estimates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_actionmaps_estimates", "SWAGGER_OVERRIDE_/api/v2/journey/actionmaps/estimates")
	journey_actionmaps_estimatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_actionmaps_estimates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_actionmaps_estimatesCmd)
}

func Cmdjourney_actionmaps_estimates() *cobra.Command {
	return journey_actionmaps_estimatesCmd
}
