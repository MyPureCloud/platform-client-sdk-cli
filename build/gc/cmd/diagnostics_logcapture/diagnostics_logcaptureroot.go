package diagnostics_logcapture

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/diagnostics_logcapture_browser"
)

func init() {
	diagnostics_logcaptureCmd.AddCommand(diagnostics_logcapture_browser.Cmddiagnostics_logcapture_browser())
	diagnostics_logcaptureCmd.Short = utils.GenerateCustomDescription(diagnostics_logcaptureCmd.Short, diagnostics_logcapture_browser.Description, )
	diagnostics_logcaptureCmd.Long = diagnostics_logcaptureCmd.Short
}
