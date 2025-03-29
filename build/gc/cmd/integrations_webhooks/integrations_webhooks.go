package integrations_webhooks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_webhooks", "SWAGGER_OVERRIDE_/api/v2/integrations/webhooks")
	integrations_webhooksCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_webhooks"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_webhooksCmd)
}

func Cmdintegrations_webhooks() *cobra.Command {
	return integrations_webhooksCmd
}
