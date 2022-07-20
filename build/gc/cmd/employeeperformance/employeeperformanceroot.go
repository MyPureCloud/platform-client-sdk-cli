package employeeperformance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/employeeperformance_externalmetrics"
)

func init() {
	employeeperformanceCmd.AddCommand(employeeperformance_externalmetrics.Cmdemployeeperformance_externalmetrics())
	employeeperformanceCmd.Short = utils.GenerateCustomDescription(employeeperformanceCmd.Short, employeeperformance_externalmetrics.Description, )
	employeeperformanceCmd.Long = employeeperformanceCmd.Short
}
