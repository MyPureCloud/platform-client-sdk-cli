package taskmanagement_worktypes_flows_onattributechange

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("taskmanagement_worktypes_flows_onattributechange", "SWAGGER_OVERRIDE_/api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange")
	taskmanagement_worktypes_flows_onattributechangeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("taskmanagement_worktypes_flows_onattributechange"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(taskmanagement_worktypes_flows_onattributechangeCmd)
}

func Cmdtaskmanagement_worktypes_flows_onattributechange() *cobra.Command {
	return taskmanagement_worktypes_flows_onattributechangeCmd
}
