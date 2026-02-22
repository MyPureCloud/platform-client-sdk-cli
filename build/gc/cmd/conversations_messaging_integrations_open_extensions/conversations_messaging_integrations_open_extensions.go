package conversations_messaging_integrations_open_extensions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("conversations_messaging_integrations_open_extensions", "SWAGGER_OVERRIDE_/api/v2/conversations/messaging/integrations/open/extensions")
	conversations_messaging_integrations_open_extensionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("conversations_messaging_integrations_open_extensions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(conversations_messaging_integrations_open_extensionsCmd)
}

func Cmdconversations_messaging_integrations_open_extensions() *cobra.Command {
	return conversations_messaging_integrations_open_extensionsCmd
}
