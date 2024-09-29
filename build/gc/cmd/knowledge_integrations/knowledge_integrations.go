package knowledge_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("knowledge_integrations", "SWAGGER_OVERRIDE_/api/v2/knowledge/integrations")
	knowledge_integrationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("knowledge_integrations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(knowledge_integrationsCmd)
}

func Cmdknowledge_integrations() *cobra.Command {
	return knowledge_integrationsCmd
}
