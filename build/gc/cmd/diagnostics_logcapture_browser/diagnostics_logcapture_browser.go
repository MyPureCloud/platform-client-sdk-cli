package diagnostics_logcapture_browser

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("diagnostics_logcapture_browser", "SWAGGER_OVERRIDE_/api/v2/diagnostics/logcapture/browser")
	diagnostics_logcapture_browserCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("diagnostics_logcapture_browser"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(diagnostics_logcapture_browserCmd)
}

func Cmddiagnostics_logcapture_browser() *cobra.Command {
	return diagnostics_logcapture_browserCmd
}
