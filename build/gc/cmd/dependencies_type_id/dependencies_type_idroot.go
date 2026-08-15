package dependencies_type_id

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dependencies_type_id_connections"
)

func init() {
	dependencies_type_idCmd.AddCommand(dependencies_type_id_connections.Cmddependencies_type_id_connections())
	dependencies_type_idCmd.Short = utils.GenerateCustomDescription(dependencies_type_idCmd.Short, dependencies_type_id_connections.Description, )
	dependencies_type_idCmd.Long = dependencies_type_idCmd.Short
}
