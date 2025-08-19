package analytics_taskmanagement_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_taskmanagement_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_taskmanagement_aggregates_jobs"
)

func init() {
	analytics_taskmanagement_aggregatesCmd.AddCommand(analytics_taskmanagement_aggregates_query.Cmdanalytics_taskmanagement_aggregates_query())
	analytics_taskmanagement_aggregatesCmd.AddCommand(analytics_taskmanagement_aggregates_jobs.Cmdanalytics_taskmanagement_aggregates_jobs())
	analytics_taskmanagement_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_taskmanagement_aggregatesCmd.Short, analytics_taskmanagement_aggregates_query.Description, analytics_taskmanagement_aggregates_jobs.Description, )
	analytics_taskmanagement_aggregatesCmd.Long = analytics_taskmanagement_aggregatesCmd.Short
}
