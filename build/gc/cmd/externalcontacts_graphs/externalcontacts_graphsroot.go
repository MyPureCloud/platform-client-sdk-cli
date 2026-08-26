package externalcontacts_graphs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_graphs_clusterscans"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/externalcontacts_graphs_settings"
)

func init() {
	externalcontacts_graphsCmd.AddCommand(externalcontacts_graphs_clusterscans.Cmdexternalcontacts_graphs_clusterscans())
	externalcontacts_graphsCmd.AddCommand(externalcontacts_graphs_settings.Cmdexternalcontacts_graphs_settings())
	externalcontacts_graphsCmd.Short = utils.GenerateCustomDescription(externalcontacts_graphsCmd.Short, externalcontacts_graphs_clusterscans.Description, externalcontacts_graphs_settings.Description, )
	externalcontacts_graphsCmd.Long = externalcontacts_graphsCmd.Short
}
