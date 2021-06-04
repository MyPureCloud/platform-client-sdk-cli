package configuration_schemas

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/configuration_schemas_edges"
)

func init() {
	configuration_schemasCmd.AddCommand(configuration_schemas_edges.Cmdconfiguration_schemas_edges())
	configuration_schemasCmd.Short = utils.GenerateCustomDescription(configuration_schemasCmd.Short, configuration_schemas_edges.Description, )
	configuration_schemasCmd.Long = configuration_schemasCmd.Short
}
