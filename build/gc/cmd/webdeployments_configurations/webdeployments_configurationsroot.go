package webdeployments_configurations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_configurations_versions"
)

func init() {
	webdeployments_configurationsCmd.AddCommand(webdeployments_configurations_versions.Cmdwebdeployments_configurations_versions())
	webdeployments_configurationsCmd.Short = utils.GenerateCustomDescription(webdeployments_configurationsCmd.Short, webdeployments_configurations_versions.Description, )
	webdeployments_configurationsCmd.Long = webdeployments_configurationsCmd.Short
}
