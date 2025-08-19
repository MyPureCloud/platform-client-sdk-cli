package taskmanagement_workitems_bulk_jobs_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_jobs_users_me"
)

func init() {
	taskmanagement_workitems_bulk_jobs_usersCmd.AddCommand(taskmanagement_workitems_bulk_jobs_users_me.Cmdtaskmanagement_workitems_bulk_jobs_users_me())
	taskmanagement_workitems_bulk_jobs_usersCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_bulk_jobs_usersCmd.Short, taskmanagement_workitems_bulk_jobs_users_me.Description, )
	taskmanagement_workitems_bulk_jobs_usersCmd.Long = taskmanagement_workitems_bulk_jobs_usersCmd.Short
}
