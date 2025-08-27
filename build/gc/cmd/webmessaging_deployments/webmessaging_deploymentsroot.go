package webmessaging_deployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webmessaging_deployments_pushdevices"
)

func init() {
	webmessaging_deploymentsCmd.AddCommand(webmessaging_deployments_pushdevices.Cmdwebmessaging_deployments_pushdevices())
	webmessaging_deploymentsCmd.Short = utils.GenerateCustomDescription(webmessaging_deploymentsCmd.Short, webmessaging_deployments_pushdevices.Description, )
	webmessaging_deploymentsCmd.Long = webmessaging_deploymentsCmd.Short
}
