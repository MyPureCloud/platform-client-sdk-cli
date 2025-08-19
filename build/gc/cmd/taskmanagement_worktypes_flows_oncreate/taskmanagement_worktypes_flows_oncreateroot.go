package taskmanagement_worktypes_flows_oncreate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_flows_oncreate_rules"
)

func init() {
	taskmanagement_worktypes_flows_oncreateCmd.AddCommand(taskmanagement_worktypes_flows_oncreate_rules.Cmdtaskmanagement_worktypes_flows_oncreate_rules())
	taskmanagement_worktypes_flows_oncreateCmd.Short = utils.GenerateCustomDescription(taskmanagement_worktypes_flows_oncreateCmd.Short, taskmanagement_worktypes_flows_oncreate_rules.Description, )
	taskmanagement_worktypes_flows_oncreateCmd.Long = taskmanagement_worktypes_flows_oncreateCmd.Short
}
