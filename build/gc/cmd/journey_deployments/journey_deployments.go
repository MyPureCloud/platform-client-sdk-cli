package journey_deployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_deployments", "SWAGGER_OVERRIDE_/api/v2/journey/deployments")
	journey_deploymentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_deployments"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_deploymentsCmd)
}

func Cmdjourney_deployments() *cobra.Command {
	return journey_deploymentsCmd
}
