package dataextensions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dataextensions_coretypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dataextensions_limits"
)

func init() {
	dataextensionsCmd.AddCommand(dataextensions_coretypes.Cmddataextensions_coretypes())
	dataextensionsCmd.AddCommand(dataextensions_limits.Cmddataextensions_limits())
	dataextensionsCmd.Short = utils.GenerateCustomDescription(dataextensionsCmd.Short, dataextensions_coretypes.Description, dataextensions_limits.Description, )
	dataextensionsCmd.Long = dataextensionsCmd.Short
}
