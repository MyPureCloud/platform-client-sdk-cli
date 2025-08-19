package flows_instances_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("flows_instances_settings", "SWAGGER_OVERRIDE_/api/v2/flows/instances/settings")
	flows_instances_settingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("flows_instances_settings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(flows_instances_settingsCmd)
}

func Cmdflows_instances_settings() *cobra.Command {
	return flows_instances_settingsCmd
}
