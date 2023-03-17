package telephony_providers_edges_sites

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_sites_numberplans"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/telephony_providers_edges_sites_outboundroutes"
)

func init() {
	telephony_providers_edges_sitesCmd.AddCommand(telephony_providers_edges_sites_numberplans.Cmdtelephony_providers_edges_sites_numberplans())
	telephony_providers_edges_sitesCmd.AddCommand(telephony_providers_edges_sites_outboundroutes.Cmdtelephony_providers_edges_sites_outboundroutes())
	telephony_providers_edges_sitesCmd.Short = utils.GenerateCustomDescription(telephony_providers_edges_sitesCmd.Short, telephony_providers_edges_sites_numberplans.Description, telephony_providers_edges_sites_outboundroutes.Description, )
	telephony_providers_edges_sitesCmd.Long = telephony_providers_edges_sitesCmd.Short
}
