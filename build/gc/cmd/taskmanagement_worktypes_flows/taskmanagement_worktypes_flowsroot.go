package taskmanagement_worktypes_flows

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_flows_datebased"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_flows_onattributechange"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_flows_oncreate"
)

func init() {
	taskmanagement_worktypes_flowsCmd.AddCommand(taskmanagement_worktypes_flows_datebased.Cmdtaskmanagement_worktypes_flows_datebased())
	taskmanagement_worktypes_flowsCmd.AddCommand(taskmanagement_worktypes_flows_onattributechange.Cmdtaskmanagement_worktypes_flows_onattributechange())
	taskmanagement_worktypes_flowsCmd.AddCommand(taskmanagement_worktypes_flows_oncreate.Cmdtaskmanagement_worktypes_flows_oncreate())
	taskmanagement_worktypes_flowsCmd.Short = utils.GenerateCustomDescription(taskmanagement_worktypes_flowsCmd.Short, taskmanagement_worktypes_flows_datebased.Description, taskmanagement_worktypes_flows_onattributechange.Description, taskmanagement_worktypes_flows_oncreate.Description, )
	taskmanagement_worktypes_flowsCmd.Long = taskmanagement_worktypes_flowsCmd.Short
}
