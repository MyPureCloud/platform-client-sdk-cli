package workforcemanagement_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_integrations", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/integrations")
	workforcemanagement_integrationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_integrations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_integrationsCmd)
}

func Cmdworkforcemanagement_integrations() *cobra.Command {
	return workforcemanagement_integrationsCmd
}
