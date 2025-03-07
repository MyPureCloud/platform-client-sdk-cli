package conversations_messaging_identityresolution_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messaging_identityresolution_integrations", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging/identityresolution/integrations")
	conversations_messaging_identityresolution_integrationsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messaging_identityresolution_integrations"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messaging_identityresolution_integrationsCmd)
}

func Cmdconversations_messaging_identityresolution_integrations() *cobra.Command {
	return conversations_messaging_identityresolution_integrationsCmd
}
