package taskmanagement_worktypes_flows_datebased

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_flows_datebased_rules"
)

func init() {
	taskmanagement_worktypes_flows_datebasedCmd.AddCommand(taskmanagement_worktypes_flows_datebased_rules.Cmdtaskmanagement_worktypes_flows_datebased_rules())
	taskmanagement_worktypes_flows_datebasedCmd.Short = utils.GenerateCustomDescription(taskmanagement_worktypes_flows_datebasedCmd.Short, taskmanagement_worktypes_flows_datebased_rules.Description, )
	taskmanagement_worktypes_flows_datebasedCmd.Long = taskmanagement_worktypes_flows_datebasedCmd.Short
}
