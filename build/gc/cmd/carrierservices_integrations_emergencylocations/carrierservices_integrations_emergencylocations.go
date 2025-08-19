package carrierservices_integrations_emergencylocations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("carrierservices_integrations_emergencylocations", "SWAGGER_OVERRIDE_/api/v2/carrierservices/integrations/emergencylocations")
	carrierservices_integrations_emergencylocationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("carrierservices_integrations_emergencylocations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(carrierservices_integrations_emergencylocationsCmd)
}

func Cmdcarrierservices_integrations_emergencylocations() *cobra.Command {
	return carrierservices_integrations_emergencylocationsCmd
}
