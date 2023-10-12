package taskmanagement_workbins

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workbins_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workbins_versions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workbins_history"
)

func init() {
	taskmanagement_workbinsCmd.AddCommand(taskmanagement_workbins_query.Cmdtaskmanagement_workbins_query())
	taskmanagement_workbinsCmd.AddCommand(taskmanagement_workbins_versions.Cmdtaskmanagement_workbins_versions())
	taskmanagement_workbinsCmd.AddCommand(taskmanagement_workbins_history.Cmdtaskmanagement_workbins_history())
	taskmanagement_workbinsCmd.Short = utils.GenerateCustomDescription(taskmanagement_workbinsCmd.Short, taskmanagement_workbins_query.Description, taskmanagement_workbins_versions.Description, taskmanagement_workbins_history.Description, )
	taskmanagement_workbinsCmd.Long = taskmanagement_workbinsCmd.Short
}
