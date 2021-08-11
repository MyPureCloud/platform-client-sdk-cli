package webdeployments_configurations_versions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/webdeployments_configurations_versions_draft"
)

func init() {
	webdeployments_configurations_versionsCmd.AddCommand(webdeployments_configurations_versions_draft.Cmdwebdeployments_configurations_versions_draft())
	webdeployments_configurations_versionsCmd.Short = utils.GenerateCustomDescription(webdeployments_configurations_versionsCmd.Short, webdeployments_configurations_versions_draft.Description, )
	webdeployments_configurations_versionsCmd.Long = webdeployments_configurations_versionsCmd.Short
}
