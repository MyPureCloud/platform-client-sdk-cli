package taskmanagement_workitems_bulk_add

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_add_jobs"
)

func init() {
	taskmanagement_workitems_bulk_addCmd.AddCommand(taskmanagement_workitems_bulk_add_jobs.Cmdtaskmanagement_workitems_bulk_add_jobs())
	taskmanagement_workitems_bulk_addCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_bulk_addCmd.Short, taskmanagement_workitems_bulk_add_jobs.Description, )
	taskmanagement_workitems_bulk_addCmd.Long = taskmanagement_workitems_bulk_addCmd.Short
}
