package socialmedia_analytics_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/socialmedia_analytics_aggregates_jobs_results"
)

func init() {
	socialmedia_analytics_aggregates_jobsCmd.AddCommand(socialmedia_analytics_aggregates_jobs_results.Cmdsocialmedia_analytics_aggregates_jobs_results())
	socialmedia_analytics_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(socialmedia_analytics_aggregates_jobsCmd.Short, socialmedia_analytics_aggregates_jobs_results.Description, )
	socialmedia_analytics_aggregates_jobsCmd.Long = socialmedia_analytics_aggregates_jobsCmd.Short
}
