package diagnostics_logcapture_browser

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/diagnostics_logcapture_browser_entries"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/diagnostics_logcapture_browser_users"
)

func init() {
	diagnostics_logcapture_browserCmd.AddCommand(diagnostics_logcapture_browser_entries.Cmddiagnostics_logcapture_browser_entries())
	diagnostics_logcapture_browserCmd.AddCommand(diagnostics_logcapture_browser_users.Cmddiagnostics_logcapture_browser_users())
	diagnostics_logcapture_browserCmd.Short = utils.GenerateCustomDescription(diagnostics_logcapture_browserCmd.Short, diagnostics_logcapture_browser_entries.Description, diagnostics_logcapture_browser_users.Description, )
	diagnostics_logcapture_browserCmd.Long = diagnostics_logcapture_browserCmd.Short
}
