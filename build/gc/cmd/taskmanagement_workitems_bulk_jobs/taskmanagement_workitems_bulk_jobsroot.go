package taskmanagement_workitems_bulk_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_jobs_users"
)

func init() {
	taskmanagement_workitems_bulk_jobsCmd.AddCommand(taskmanagement_workitems_bulk_jobs_users.Cmdtaskmanagement_workitems_bulk_jobs_users())
	taskmanagement_workitems_bulk_jobsCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_bulk_jobsCmd.Short, taskmanagement_workitems_bulk_jobs_users.Description, )
	taskmanagement_workitems_bulk_jobsCmd.Long = taskmanagement_workitems_bulk_jobsCmd.Short
}
