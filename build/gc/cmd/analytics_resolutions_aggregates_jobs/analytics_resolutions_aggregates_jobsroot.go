package analytics_resolutions_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_resolutions_aggregates_jobs_results"
)

func init() {
	analytics_resolutions_aggregates_jobsCmd.AddCommand(analytics_resolutions_aggregates_jobs_results.Cmdanalytics_resolutions_aggregates_jobs_results())
	analytics_resolutions_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_resolutions_aggregates_jobsCmd.Short, analytics_resolutions_aggregates_jobs_results.Description, )
	analytics_resolutions_aggregates_jobsCmd.Long = analytics_resolutions_aggregates_jobsCmd.Short
}
