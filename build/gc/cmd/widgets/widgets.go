package widgets

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("widgets", "SWAGGER_OVERRIDE_/api/v2/widgets")
	widgetsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("widgets"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(widgetsCmd)
}

func Cmdwidgets() *cobra.Command {
	return widgetsCmd
}
