package employeeperformance_externalmetrics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/employeeperformance_externalmetrics_data"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/employeeperformance_externalmetrics_definitions"
)

func init() {
	employeeperformance_externalmetricsCmd.AddCommand(employeeperformance_externalmetrics_data.Cmdemployeeperformance_externalmetrics_data())
	employeeperformance_externalmetricsCmd.AddCommand(employeeperformance_externalmetrics_definitions.Cmdemployeeperformance_externalmetrics_definitions())
	employeeperformance_externalmetricsCmd.Short = utils.GenerateCustomDescription(employeeperformance_externalmetricsCmd.Short, employeeperformance_externalmetrics_data.Description, employeeperformance_externalmetrics_definitions.Description, )
	employeeperformance_externalmetricsCmd.Long = employeeperformance_externalmetricsCmd.Short
}
