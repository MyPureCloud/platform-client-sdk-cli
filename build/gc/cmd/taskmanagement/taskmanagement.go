package taskmanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement", "SWAGGER_OVERRIDE_/api/v2/taskmanagement")
	taskmanagementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagementCmd)
}

func Cmdtaskmanagement() *cobra.Command {
	return taskmanagementCmd
}
