package taskmanagement_worktypes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_statuses"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_history"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes_flows"
)

func init() {
	taskmanagement_worktypesCmd.AddCommand(taskmanagement_worktypes_statuses.Cmdtaskmanagement_worktypes_statuses())
	taskmanagement_worktypesCmd.AddCommand(taskmanagement_worktypes_query.Cmdtaskmanagement_worktypes_query())
	taskmanagement_worktypesCmd.AddCommand(taskmanagement_worktypes_versions.Cmdtaskmanagement_worktypes_versions())
	taskmanagement_worktypesCmd.AddCommand(taskmanagement_worktypes_history.Cmdtaskmanagement_worktypes_history())
	taskmanagement_worktypesCmd.AddCommand(taskmanagement_worktypes_flows.Cmdtaskmanagement_worktypes_flows())
	taskmanagement_worktypesCmd.Short = utils.GenerateCustomDescription(taskmanagement_worktypesCmd.Short, taskmanagement_worktypes_statuses.Description, taskmanagement_worktypes_query.Description, taskmanagement_worktypes_versions.Description, taskmanagement_worktypes_history.Description, taskmanagement_worktypes_flows.Description, )
	taskmanagement_worktypesCmd.Long = taskmanagement_worktypesCmd.Short
}
