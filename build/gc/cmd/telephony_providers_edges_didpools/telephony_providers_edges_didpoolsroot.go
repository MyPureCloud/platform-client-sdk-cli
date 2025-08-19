package telephony_providers_edges_didpools

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_didpools_dids"
)

func init() {
	telephony_providers_edges_didpoolsCmd.AddCommand(telephony_providers_edges_didpools_dids.Cmdtelephony_providers_edges_didpools_dids())
	telephony_providers_edges_didpoolsCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_didpoolsCmd.Short, telephony_providers_edges_didpools_dids.Description, )
	telephony_providers_edges_didpoolsCmd.Long = telephony_providers_edges_didpoolsCmd.Short
}
