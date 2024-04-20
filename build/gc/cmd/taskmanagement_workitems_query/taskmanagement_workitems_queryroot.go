package taskmanagement_workitems_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/taskmanagement_workitems_query_jobs"
)

func init() {
	taskmanagement_workitems_queryCmd.AddCommand(taskmanagement_workitems_query_jobs.Cmdtaskmanagement_workitems_query_jobs())
	taskmanagement_workitems_queryCmd.Short = utils.GenerateCustomDescription(taskmanagement_workitems_queryCmd.Short, taskmanagement_workitems_query_jobs.Description, )
	taskmanagement_workitems_queryCmd.Long = taskmanagement_workitems_queryCmd.Short
}
