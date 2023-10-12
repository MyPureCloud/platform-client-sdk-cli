package analytics_actions_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_actions_aggregates_jobs_results"
)

func init() {
	analytics_actions_aggregates_jobsCmd.AddCommand(analytics_actions_aggregates_jobs_results.Cmdanalytics_actions_aggregates_jobs_results())
	analytics_actions_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_actions_aggregates_jobsCmd.Short, analytics_actions_aggregates_jobs_results.Description, )
	analytics_actions_aggregates_jobsCmd.Long = analytics_actions_aggregates_jobsCmd.Short
}
