package workforcemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement")
	workforcemanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagementCmd)
}

func Cmdworkforcemanagement() *cobra.Command {
	return workforcemanagementCmd
}
