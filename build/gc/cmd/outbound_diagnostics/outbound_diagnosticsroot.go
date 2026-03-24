package outbound_diagnostics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_diagnostics_campaigns"
)

func init() {
	outbound_diagnosticsCmd.AddCommand(outbound_diagnostics_campaigns.Cmdoutbound_diagnostics_campaigns())
	outbound_diagnosticsCmd.Short = utils.GenerateCustomDescription(outbound_diagnosticsCmd.Short, outbound_diagnostics_campaigns.Description, )
	outbound_diagnosticsCmd.Long = outbound_diagnosticsCmd.Short
}
