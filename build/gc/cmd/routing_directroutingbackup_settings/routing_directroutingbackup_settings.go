package routing_directroutingbackup_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("routing_directroutingbackup_settings", "SWAGGER_OVERRIDE_/api/v2/routing/directroutingbackup/settings")
	routing_directroutingbackup_settingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("routing_directroutingbackup_settings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(routing_directroutingbackup_settingsCmd)
}

func Cmdrouting_directroutingbackup_settings() *cobra.Command {
	return routing_directroutingbackup_settingsCmd
}
