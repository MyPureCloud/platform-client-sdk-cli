package taskmanagement_workitems_bulk_add_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_add_jobs_results"
)

func init() {
	taskmanagement_workitems_bulk_add_jobsCmd.AddCommand(taskmanagement_workitems_bulk_add_jobs_results.Cmdtaskmanagement_workitems_bulk_add_jobs_results())
	taskmanagement_workitems_bulk_add_jobsCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_bulk_add_jobsCmd.Short, taskmanagement_workitems_bulk_add_jobs_results.Description, )
	taskmanagement_workitems_bulk_add_jobsCmd.Long = taskmanagement_workitems_bulk_add_jobsCmd.Short
}
