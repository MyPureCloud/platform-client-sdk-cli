package dependencies

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dependencies_type"
)

func init() {
	dependenciesCmd.AddCommand(dependencies_type.Cmddependencies_type())
	dependenciesCmd.Short = utils.GenerateCustomDescription(dependenciesCmd.Short, dependencies_type.Description, )
	dependenciesCmd.Long = dependenciesCmd.Short
}
