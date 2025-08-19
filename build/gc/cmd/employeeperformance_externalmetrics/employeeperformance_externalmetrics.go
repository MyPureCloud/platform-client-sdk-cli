package employeeperformance_externalmetrics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("employeeperformance_externalmetrics", "SWAGGER_OVERRIDE_/api/v2/employeeperformance/externalmetrics")
	employeeperformance_externalmetricsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("employeeperformance_externalmetrics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(employeeperformance_externalmetricsCmd)
}

func Cmdemployeeperformance_externalmetrics() *cobra.Command {
	return employeeperformance_externalmetricsCmd
}
