package analytics_copilots_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_copilots_aggregates_jobs_results"
)

func init() {
	analytics_copilots_aggregates_jobsCmd.AddCommand(analytics_copilots_aggregates_jobs_results.Cmdanalytics_copilots_aggregates_jobs_results())
	analytics_copilots_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_copilots_aggregates_jobsCmd.Short, analytics_copilots_aggregates_jobs_results.Description, )
	analytics_copilots_aggregates_jobsCmd.Long = analytics_copilots_aggregates_jobsCmd.Short
}
