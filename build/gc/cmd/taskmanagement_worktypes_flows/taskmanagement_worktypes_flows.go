package taskmanagement_worktypes_flows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_worktypes_flows", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows")
	taskmanagement_worktypes_flowsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_worktypes_flows"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_worktypes_flowsCmd)
}

func Cmdtaskmanagement_worktypes_flows() *cobra.Command {
	return taskmanagement_worktypes_flowsCmd
}
