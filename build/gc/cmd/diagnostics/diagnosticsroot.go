package diagnostics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/diagnostics_logcapture"
)

func init() {
	diagnosticsCmd.AddCommand(diagnostics_logcapture.Cmddiagnostics_logcapture())
	diagnosticsCmd.Short = utils.GenerateCustomDescription(diagnosticsCmd.Short, diagnostics_logcapture.Description, )
	diagnosticsCmd.Long = diagnosticsCmd.Short
}
