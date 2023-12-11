package diagnostics_logcapture_browser_entries_download

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("diagnostics_logcapture_browser_entries_download", "SWAGGER_OVERRIDE_/api/v2/diagnostics/logcapture/browser/entries/download")
	diagnostics_logcapture_browser_entries_downloadCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("diagnostics_logcapture_browser_entries_download"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(diagnostics_logcapture_browser_entries_downloadCmd)
}

func Cmddiagnostics_logcapture_browser_entries_download() *cobra.Command {
	return diagnostics_logcapture_browser_entries_downloadCmd
}
