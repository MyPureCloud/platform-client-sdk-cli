package taskmanagement_workitems_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_users_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_users_wrapups"
)

func init() {
	taskmanagement_workitems_usersCmd.AddCommand(taskmanagement_workitems_users_me.Cmdtaskmanagement_workitems_users_me())
	taskmanagement_workitems_usersCmd.AddCommand(taskmanagement_workitems_users_wrapups.Cmdtaskmanagement_workitems_users_wrapups())
	taskmanagement_workitems_usersCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_usersCmd.Short, taskmanagement_workitems_users_me.Description, taskmanagement_workitems_users_wrapups.Description, )
	taskmanagement_workitems_usersCmd.Long = taskmanagement_workitems_usersCmd.Short
}
