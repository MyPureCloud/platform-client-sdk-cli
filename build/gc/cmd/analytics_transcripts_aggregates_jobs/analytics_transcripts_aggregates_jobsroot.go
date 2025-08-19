package analytics_transcripts_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_transcripts_aggregates_jobs_results"
)

func init() {
	analytics_transcripts_aggregates_jobsCmd.AddCommand(analytics_transcripts_aggregates_jobs_results.Cmdanalytics_transcripts_aggregates_jobs_results())
	analytics_transcripts_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_transcripts_aggregates_jobsCmd.Short, analytics_transcripts_aggregates_jobs_results.Description, )
	analytics_transcripts_aggregates_jobsCmd.Long = analytics_transcripts_aggregates_jobsCmd.Short
}
