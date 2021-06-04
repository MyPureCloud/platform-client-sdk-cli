package configuration

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/configuration_schemas"
)

func init() {
	configurationCmd.AddCommand(configuration_schemas.Cmdconfiguration_schemas())
	configurationCmd.Short = utils.GenerateCustomDescription(configurationCmd.Short, configuration_schemas.Description, )
	configurationCmd.Long = configurationCmd.Short
}
