package integrations_config

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("integrations_config", "SWAGGER_OVERRIDE_/api/v2/integrations/{integrationId}/config")
	integrations_configCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("integrations_config"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(integrations_configCmd)
}

func Cmdintegrations_config() *cobra.Command {
	return integrations_configCmd
}
