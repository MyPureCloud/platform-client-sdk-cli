package analytics_bots_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_bots_aggregates_jobs_results"
)

func init() {
	analytics_bots_aggregates_jobsCmd.AddCommand(analytics_bots_aggregates_jobs_results.Cmdanalytics_bots_aggregates_jobs_results())
	analytics_bots_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_bots_aggregates_jobsCmd.Short, analytics_bots_aggregates_jobs_results.Description, )
	analytics_bots_aggregates_jobsCmd.Long = analytics_bots_aggregates_jobsCmd.Short
}
