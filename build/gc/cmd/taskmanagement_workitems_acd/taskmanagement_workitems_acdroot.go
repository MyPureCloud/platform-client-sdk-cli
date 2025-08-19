package taskmanagement_workitems_acd

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_acd_cancel"
)

func init() {
	taskmanagement_workitems_acdCmd.AddCommand(taskmanagement_workitems_acd_cancel.Cmdtaskmanagement_workitems_acd_cancel())
	taskmanagement_workitems_acdCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_acdCmd.Short, taskmanagement_workitems_acd_cancel.Description, )
	taskmanagement_workitems_acdCmd.Long = taskmanagement_workitems_acdCmd.Short
}
