package telephony_providers_edges_trunks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_trunks_metrics"
)

func init() {
	telephony_providers_edges_trunksCmd.AddCommand(telephony_providers_edges_trunks_metrics.Cmdtelephony_providers_edges_trunks_metrics())
	telephony_providers_edges_trunksCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_trunksCmd.Short, telephony_providers_edges_trunks_metrics.Description, )
	telephony_providers_edges_trunksCmd.Long = telephony_providers_edges_trunksCmd.Short
}
