package flows_instances_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_instances_settings_executiondata"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_instances_settings_loglevels"
)

func init() {
	flows_instances_settingsCmd.AddCommand(flows_instances_settings_executiondata.Cmdflows_instances_settings_executiondata())
	flows_instances_settingsCmd.AddCommand(flows_instances_settings_loglevels.Cmdflows_instances_settings_loglevels())
	flows_instances_settingsCmd.Short = utils.GenerateCustomDescription(flows_instances_settingsCmd.Short, flows_instances_settings_executiondata.Description, flows_instances_settings_loglevels.Description, )
	flows_instances_settingsCmd.Long = flows_instances_settingsCmd.Short
}
