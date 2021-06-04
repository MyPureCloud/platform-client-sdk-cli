package configuration_schemas_edges

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/configuration_schemas_edges_vnext"
)

func init() {
	configuration_schemas_edgesCmd.AddCommand(configuration_schemas_edges_vnext.Cmdconfiguration_schemas_edges_vnext())
	configuration_schemas_edgesCmd.Short = utils.GenerateCustomDescription(configuration_schemas_edgesCmd.Short, configuration_schemas_edges_vnext.Description, )
	configuration_schemas_edgesCmd.Long = configuration_schemas_edgesCmd.Short
}
