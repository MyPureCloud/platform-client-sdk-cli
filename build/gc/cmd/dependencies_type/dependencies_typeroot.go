package dependencies_type

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dependencies_type_id"
)

func init() {
	dependencies_typeCmd.AddCommand(dependencies_type_id.Cmddependencies_type_id())
	dependencies_typeCmd.Short = utils.GenerateCustomDescription(dependencies_typeCmd.Short, dependencies_type_id.Description, )
	dependencies_typeCmd.Long = dependencies_typeCmd.Short
}
