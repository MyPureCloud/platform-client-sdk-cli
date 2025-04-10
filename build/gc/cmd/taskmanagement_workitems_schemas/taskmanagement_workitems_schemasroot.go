package taskmanagement_workitems_schemas

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_schemas_coretypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_schemas_limits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_schemas_versions"
)

func init() {
	taskmanagement_workitems_schemasCmd.AddCommand(taskmanagement_workitems_schemas_coretypes.Cmdtaskmanagement_workitems_schemas_coretypes())
	taskmanagement_workitems_schemasCmd.AddCommand(taskmanagement_workitems_schemas_limits.Cmdtaskmanagement_workitems_schemas_limits())
	taskmanagement_workitems_schemasCmd.AddCommand(taskmanagement_workitems_schemas_versions.Cmdtaskmanagement_workitems_schemas_versions())
	taskmanagement_workitems_schemasCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_schemasCmd.Short, taskmanagement_workitems_schemas_coretypes.Description, taskmanagement_workitems_schemas_limits.Description, taskmanagement_workitems_schemas_versions.Description, )
	taskmanagement_workitems_schemasCmd.Long = taskmanagement_workitems_schemasCmd.Short
}
