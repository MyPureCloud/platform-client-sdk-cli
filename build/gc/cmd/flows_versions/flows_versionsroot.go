package flows_versions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_versions_intents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_versions_health"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_versions_configuration"
)

func init() {
	flows_versionsCmd.AddCommand(flows_versions_intents.Cmdflows_versions_intents())
	flows_versionsCmd.AddCommand(flows_versions_health.Cmdflows_versions_health())
	flows_versionsCmd.AddCommand(flows_versions_configuration.Cmdflows_versions_configuration())
	flows_versionsCmd.Short = utils.GenerateCustomDescription(flows_versionsCmd.Short, flows_versions_intents.Description, flows_versions_health.Description, flows_versions_configuration.Description, )
	flows_versionsCmd.Long = flows_versionsCmd.Short
}
