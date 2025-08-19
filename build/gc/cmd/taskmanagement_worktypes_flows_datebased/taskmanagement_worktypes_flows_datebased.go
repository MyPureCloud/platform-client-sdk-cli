package taskmanagement_worktypes_flows_datebased

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_worktypes_flows_datebased", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased")
	taskmanagement_worktypes_flows_datebasedCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_worktypes_flows_datebased"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_worktypes_flows_datebasedCmd)
}

func Cmdtaskmanagement_worktypes_flows_datebased() *cobra.Command {
	return taskmanagement_worktypes_flows_datebasedCmd
}
