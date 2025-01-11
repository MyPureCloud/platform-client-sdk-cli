package taskmanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workbins"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_worktypes"
)

func init() {
	taskmanagementCmd.AddCommand(taskmanagement_workitems.Cmdtaskmanagement_workitems())
	taskmanagementCmd.AddCommand(taskmanagement_workbins.Cmdtaskmanagement_workbins())
	taskmanagementCmd.AddCommand(taskmanagement_worktypes.Cmdtaskmanagement_worktypes())
	taskmanagementCmd.Short = utils.GenerateCustomDescription(taskmanagementCmd.Short, taskmanagement_workitems.Description, taskmanagement_workbins.Description, taskmanagement_worktypes.Description, )
	taskmanagementCmd.Long = taskmanagementCmd.Short
}
