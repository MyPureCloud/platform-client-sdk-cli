package webdeployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_configurations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_deployments"
)

func init() {
	webdeploymentsCmd.AddCommand(webdeployments_configurations.Cmdwebdeployments_configurations())
	webdeploymentsCmd.AddCommand(webdeployments_deployments.Cmdwebdeployments_deployments())
	webdeploymentsCmd.Short = utils.GenerateCustomDescription(webdeploymentsCmd.Short, webdeployments_configurations.Description, webdeployments_deployments.Description, )
	webdeploymentsCmd.Long = webdeploymentsCmd.Short
}
