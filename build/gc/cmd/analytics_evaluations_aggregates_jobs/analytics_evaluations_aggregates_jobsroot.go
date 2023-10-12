package analytics_evaluations_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_evaluations_aggregates_jobs_results"
)

func init() {
	analytics_evaluations_aggregates_jobsCmd.AddCommand(analytics_evaluations_aggregates_jobs_results.Cmdanalytics_evaluations_aggregates_jobs_results())
	analytics_evaluations_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_evaluations_aggregates_jobsCmd.Short, analytics_evaluations_aggregates_jobs_results.Description, )
	analytics_evaluations_aggregates_jobsCmd.Long = analytics_evaluations_aggregates_jobsCmd.Short
}
