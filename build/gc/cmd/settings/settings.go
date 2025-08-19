package settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("settings", "SWAGGER_OVERRIDE_/api/v2/settings")
	settingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("settings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(settingsCmd)
}

func Cmdsettings() *cobra.Command {
	return settingsCmd
}
