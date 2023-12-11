package diagnostics_logcapture_browser_entries

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/diagnostics_logcapture_browser_entries_download"
)

func init() {
	diagnostics_logcapture_browser_entriesCmd.AddCommand(diagnostics_logcapture_browser_entries_download.Cmddiagnostics_logcapture_browser_entries_download())
	diagnostics_logcapture_browser_entriesCmd.Short = utils.GenerateCustomDescription(diagnostics_logcapture_browser_entriesCmd.Short, diagnostics_logcapture_browser_entries_download.Description, )
	diagnostics_logcapture_browser_entriesCmd.Long = diagnostics_logcapture_browser_entriesCmd.Short
}
