package analytics_taskmanagement_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_taskmanagement_aggregates_jobs_results"
)

func init() {
	analytics_taskmanagement_aggregates_jobsCmd.AddCommand(analytics_taskmanagement_aggregates_jobs_results.Cmdanalytics_taskmanagement_aggregates_jobs_results())
	analytics_taskmanagement_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_taskmanagement_aggregates_jobsCmd.Short, analytics_taskmanagement_aggregates_jobs_results.Description, )
	analytics_taskmanagement_aggregates_jobsCmd.Long = analytics_taskmanagement_aggregates_jobsCmd.Short
}
