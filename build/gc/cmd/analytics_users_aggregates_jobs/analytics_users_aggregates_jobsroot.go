package analytics_users_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_users_aggregates_jobs_results"
)

func init() {
	analytics_users_aggregates_jobsCmd.AddCommand(analytics_users_aggregates_jobs_results.Cmdanalytics_users_aggregates_jobs_results())
	analytics_users_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_users_aggregates_jobsCmd.Short, analytics_users_aggregates_jobs_results.Description, )
	analytics_users_aggregates_jobsCmd.Long = analytics_users_aggregates_jobsCmd.Short
}
