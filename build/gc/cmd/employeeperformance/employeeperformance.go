package employeeperformance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("employeeperformance", "SWAGGER_OVERRIDE_/api/v2/employeeperformance")
	employeeperformanceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("employeeperformance"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(employeeperformanceCmd)
}

func Cmdemployeeperformance() *cobra.Command {
	return employeeperformanceCmd
}
