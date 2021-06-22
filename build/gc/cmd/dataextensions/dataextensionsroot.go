package dataextensions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dataextensions_limits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/dataextensions_coretypes"
)

func init() {
	dataextensionsCmd.AddCommand(dataextensions_limits.Cmddataextensions_limits())
	dataextensionsCmd.AddCommand(dataextensions_coretypes.Cmddataextensions_coretypes())
	dataextensionsCmd.Short = utils.GenerateCustomDescription(dataextensionsCmd.Short, dataextensions_limits.Description, dataextensions_coretypes.Description, )
	dataextensionsCmd.Long = dataextensionsCmd.Short
}
