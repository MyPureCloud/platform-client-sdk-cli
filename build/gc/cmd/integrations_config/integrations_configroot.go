package integrations_config

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_config_current"
)

func init() {
	integrations_configCmd.AddCommand(integrations_config_current.Cmdintegrations_config_current())
	integrations_configCmd.Short = utils.GenerateCustomDescription(integrations_configCmd.Short, integrations_config_current.Description, )
	integrations_configCmd.Long = integrations_configCmd.Short
}
