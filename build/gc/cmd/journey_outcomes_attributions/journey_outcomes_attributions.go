package journey_outcomes_attributions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_outcomes_attributions", "SWAGGER_OVERRIDE_/api/v2/journey/outcomes/attributions")
	journey_outcomes_attributionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_outcomes_attributions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_outcomes_attributionsCmd)
}

func Cmdjourney_outcomes_attributions() *cobra.Command {
	return journey_outcomes_attributionsCmd
}
