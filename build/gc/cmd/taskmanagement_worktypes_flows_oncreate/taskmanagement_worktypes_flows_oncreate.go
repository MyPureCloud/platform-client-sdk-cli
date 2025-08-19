package taskmanagement_worktypes_flows_oncreate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_worktypes_flows_oncreate", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate")
	taskmanagement_worktypes_flows_oncreateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_worktypes_flows_oncreate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_worktypes_flows_oncreateCmd)
}

func Cmdtaskmanagement_worktypes_flows_oncreate() *cobra.Command {
	return taskmanagement_worktypes_flows_oncreateCmd
}
