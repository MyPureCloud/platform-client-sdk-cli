package analytics_summaries_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_summaries_aggregates_jobs_results"
)

func init() {
	analytics_summaries_aggregates_jobsCmd.AddCommand(analytics_summaries_aggregates_jobs_results.Cmdanalytics_summaries_aggregates_jobs_results())
	analytics_summaries_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_summaries_aggregates_jobsCmd.Short, analytics_summaries_aggregates_jobs_results.Description, )
	analytics_summaries_aggregates_jobsCmd.Long = analytics_summaries_aggregates_jobsCmd.Short
}
