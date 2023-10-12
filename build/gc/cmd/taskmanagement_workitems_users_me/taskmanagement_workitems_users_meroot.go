package taskmanagement_workitems_users_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_users_me_wrapups"
)

func init() {
	taskmanagement_workitems_users_meCmd.AddCommand(taskmanagement_workitems_users_me_wrapups.Cmdtaskmanagement_workitems_users_me_wrapups())
	taskmanagement_workitems_users_meCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_users_meCmd.Short, taskmanagement_workitems_users_me_wrapups.Description, )
	taskmanagement_workitems_users_meCmd.Long = taskmanagement_workitems_users_meCmd.Short
}
