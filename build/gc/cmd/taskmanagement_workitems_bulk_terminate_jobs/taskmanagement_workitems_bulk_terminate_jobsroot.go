package taskmanagement_workitems_bulk_terminate_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_terminate_jobs_results"
)

func init() {
	taskmanagement_workitems_bulk_terminate_jobsCmd.AddCommand(taskmanagement_workitems_bulk_terminate_jobs_results.Cmdtaskmanagement_workitems_bulk_terminate_jobs_results())
	taskmanagement_workitems_bulk_terminate_jobsCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_bulk_terminate_jobsCmd.Short, taskmanagement_workitems_bulk_terminate_jobs_results.Description, )
	taskmanagement_workitems_bulk_terminate_jobsCmd.Long = taskmanagement_workitems_bulk_terminate_jobsCmd.Short
}
