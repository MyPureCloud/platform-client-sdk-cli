package carrierservices_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("carrierservices_integrations", "SWAGGER_OVERRIDE_/api/v2/carrierservices/integrations")
	carrierservices_integrationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("carrierservices_integrations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(carrierservices_integrationsCmd)
}

func Cmdcarrierservices_integrations() *cobra.Command {
	return carrierservices_integrationsCmd
}
