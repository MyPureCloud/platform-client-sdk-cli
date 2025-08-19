package diagnostics_logcapture_browser_entries

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("diagnostics_logcapture_browser_entries", "SWAGGER_OVERRIDE_/api/v2/diagnostics/logcapture/browser/entries")
	diagnostics_logcapture_browser_entriesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("diagnostics_logcapture_browser_entries"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(diagnostics_logcapture_browser_entriesCmd)
}

func Cmddiagnostics_logcapture_browser_entries() *cobra.Command {
	return diagnostics_logcapture_browser_entriesCmd
}
