package outbound_diagnostics_campaigns

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/outbound_diagnostics_campaigns_summary"
)

func init() {
	outbound_diagnostics_campaignsCmd.AddCommand(outbound_diagnostics_campaigns_summary.Cmdoutbound_diagnostics_campaigns_summary())
	outbound_diagnostics_campaignsCmd.Short = utils.GenerateCustomDescription(outbound_diagnostics_campaignsCmd.Short, outbound_diagnostics_campaigns_summary.Description, )
	outbound_diagnostics_campaignsCmd.Long = outbound_diagnostics_campaignsCmd.Short
}
