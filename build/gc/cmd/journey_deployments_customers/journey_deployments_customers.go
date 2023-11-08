package journey_deployments_customers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("journey_deployments_customers", "SWAGGER_OVERRIDE_/api/v2/journey/deployments/{deploymentId}/customers")
	journey_deployments_customersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("journey_deployments_customers"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(journey_deployments_customersCmd)
}

func Cmdjourney_deployments_customers() *cobra.Command {
	return journey_deployments_customersCmd
}
