package journey_flows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_flows", "SWAGGER_OVERRIDE_/api/v2/journey/flows")
	journey_flowsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_flows"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_flowsCmd)
}

func Cmdjourney_flows() *cobra.Command {
	return journey_flowsCmd
}
