package analytics_flowexecutions_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flowexecutions_aggregates_jobs_results"
)

func init() {
	analytics_flowexecutions_aggregates_jobsCmd.AddCommand(analytics_flowexecutions_aggregates_jobs_results.Cmdanalytics_flowexecutions_aggregates_jobs_results())
	analytics_flowexecutions_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_flowexecutions_aggregates_jobsCmd.Short, analytics_flowexecutions_aggregates_jobs_results.Description, )
	analytics_flowexecutions_aggregates_jobsCmd.Long = analytics_flowexecutions_aggregates_jobsCmd.Short
}
