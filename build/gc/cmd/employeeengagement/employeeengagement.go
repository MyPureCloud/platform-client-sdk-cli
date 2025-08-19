package employeeengagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("employeeengagement", "SWAGGER_OVERRIDE_/api/v2/employeeengagement")
	employeeengagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("employeeengagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(employeeengagementCmd)
}

func Cmdemployeeengagement() *cobra.Command {
	return employeeengagementCmd
}
