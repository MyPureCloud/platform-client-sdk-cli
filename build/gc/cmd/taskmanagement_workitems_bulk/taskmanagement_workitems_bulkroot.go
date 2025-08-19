package taskmanagement_workitems_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_add"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk_terminate"
)

func init() {
	taskmanagement_workitems_bulkCmd.AddCommand(taskmanagement_workitems_bulk_jobs.Cmdtaskmanagement_workitems_bulk_jobs())
	taskmanagement_workitems_bulkCmd.AddCommand(taskmanagement_workitems_bulk_add.Cmdtaskmanagement_workitems_bulk_add())
	taskmanagement_workitems_bulkCmd.AddCommand(taskmanagement_workitems_bulk_terminate.Cmdtaskmanagement_workitems_bulk_terminate())
	taskmanagement_workitems_bulkCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_bulkCmd.Short, taskmanagement_workitems_bulk_jobs.Description, taskmanagement_workitems_bulk_add.Description, taskmanagement_workitems_bulk_terminate.Description, )
	taskmanagement_workitems_bulkCmd.Long = taskmanagement_workitems_bulkCmd.Short
}
