package webdeployments_deployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_deployments_cobrowse"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_deployments_configurations"
)

func init() {
	webdeployments_deploymentsCmd.AddCommand(webdeployments_deployments_cobrowse.Cmdwebdeployments_deployments_cobrowse())
	webdeployments_deploymentsCmd.AddCommand(webdeployments_deployments_configurations.Cmdwebdeployments_deployments_configurations())
	webdeployments_deploymentsCmd.Short = utils.GenerateCustomDescription(webdeployments_deploymentsCmd.Short, webdeployments_deployments_cobrowse.Description, webdeployments_deployments_configurations.Description, )
	webdeployments_deploymentsCmd.Long = webdeployments_deploymentsCmd.Short
}
