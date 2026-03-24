package workforcemanagement_agents_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_integrations", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/integrations")
	workforcemanagement_agents_integrationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_integrations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_integrationsCmd)
}

func Cmdworkforcemanagement_agents_integrations() *cobra.Command {
	return workforcemanagement_agents_integrationsCmd
}
