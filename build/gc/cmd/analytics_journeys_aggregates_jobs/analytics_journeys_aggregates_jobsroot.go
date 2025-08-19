package analytics_journeys_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_journeys_aggregates_jobs_results"
)

func init() {
	analytics_journeys_aggregates_jobsCmd.AddCommand(analytics_journeys_aggregates_jobs_results.Cmdanalytics_journeys_aggregates_jobs_results())
	analytics_journeys_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_journeys_aggregates_jobsCmd.Short, analytics_journeys_aggregates_jobs_results.Description, )
	analytics_journeys_aggregates_jobsCmd.Long = analytics_journeys_aggregates_jobsCmd.Short
}
