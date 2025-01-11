package taskmanagement_workitems_bulk_terminate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_terminate_jobs"
)

func init() {
	taskmanagement_workitems_bulk_terminateCmd.AddCommand(taskmanagement_workitems_bulk_terminate_jobs.Cmdtaskmanagement_workitems_bulk_terminate_jobs())
	taskmanagement_workitems_bulk_terminateCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_bulk_terminateCmd.Short, taskmanagement_workitems_bulk_terminate_jobs.Description, )
	taskmanagement_workitems_bulk_terminateCmd.Long = taskmanagement_workitems_bulk_terminateCmd.Short
}
