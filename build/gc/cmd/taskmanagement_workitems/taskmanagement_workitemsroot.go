package taskmanagement_workitems

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_bulk"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_disconnect"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_terminate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_assignment"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_acd"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_wrapups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_history"
)

func init() {
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_bulk.Cmdtaskmanagement_workitems_bulk())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_query.Cmdtaskmanagement_workitems_query())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_disconnect.Cmdtaskmanagement_workitems_disconnect())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_terminate.Cmdtaskmanagement_workitems_terminate())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_assignment.Cmdtaskmanagement_workitems_assignment())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_acd.Cmdtaskmanagement_workitems_acd())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_schemas.Cmdtaskmanagement_workitems_schemas())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_users.Cmdtaskmanagement_workitems_users())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_wrapups.Cmdtaskmanagement_workitems_wrapups())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_versions.Cmdtaskmanagement_workitems_versions())
	taskmanagement_workitemsCmd.AddCommand(taskmanagement_workitems_history.Cmdtaskmanagement_workitems_history())
	taskmanagement_workitemsCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitemsCmd.Short, taskmanagement_workitems_bulk.Description, taskmanagement_workitems_query.Description, taskmanagement_workitems_disconnect.Description, taskmanagement_workitems_terminate.Description, taskmanagement_workitems_assignment.Description, taskmanagement_workitems_acd.Description, taskmanagement_workitems_schemas.Description, taskmanagement_workitems_users.Description, taskmanagement_workitems_wrapups.Description, taskmanagement_workitems_versions.Description, taskmanagement_workitems_history.Description, )
	taskmanagement_workitemsCmd.Long = taskmanagement_workitemsCmd.Short
}
