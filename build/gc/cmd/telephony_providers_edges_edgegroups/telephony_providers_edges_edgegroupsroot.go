package telephony_providers_edges_edgegroups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_edgegroups_edgetrunkbases"
)

func init() {
	telephony_providers_edges_edgegroupsCmd.AddCommand(telephony_providers_edges_edgegroups_edgetrunkbases.Cmdtelephony_providers_edges_edgegroups_edgetrunkbases())
	telephony_providers_edges_edgegroupsCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_edgegroupsCmd.Short, telephony_providers_edges_edgegroups_edgetrunkbases.Description, )
	telephony_providers_edges_edgegroupsCmd.Long = telephony_providers_edges_edgegroupsCmd.Short
}
