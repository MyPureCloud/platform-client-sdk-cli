package webdeployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_deployments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_token"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_configurations"
)

func init() {
	webdeploymentsCmd.AddCommand(webdeployments_deployments.Cmdwebdeployments_deployments())
	webdeploymentsCmd.AddCommand(webdeployments_token.Cmdwebdeployments_token())
	webdeploymentsCmd.AddCommand(webdeployments_configurations.Cmdwebdeployments_configurations())
	webdeploymentsCmd.Short = utils.GenerateCustomDescription(webdeploymentsCmd.Short, webdeployments_deployments.Description, webdeployments_token.Description, webdeployments_configurations.Description, )
	webdeploymentsCmd.Long = webdeploymentsCmd.Short
}
