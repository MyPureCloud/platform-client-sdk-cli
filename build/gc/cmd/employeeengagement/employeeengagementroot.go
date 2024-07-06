package employeeengagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/employeeengagement_celebrations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/employeeengagement_recognitions"
)

func init() {
	employeeengagementCmd.AddCommand(employeeengagement_celebrations.Cmdemployeeengagement_celebrations())
	employeeengagementCmd.AddCommand(employeeengagement_recognitions.Cmdemployeeengagement_recognitions())
	employeeengagementCmd.Short = utils.GenerateCustomDescription(employeeengagementCmd.Short, employeeengagement_celebrations.Description, employeeengagement_recognitions.Description, )
	employeeengagementCmd.Long = employeeengagementCmd.Short
}
