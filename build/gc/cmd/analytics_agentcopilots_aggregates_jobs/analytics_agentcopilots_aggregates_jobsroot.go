package analytics_agentcopilots_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agentcopilots_aggregates_jobs_results"
)

func init() {
	analytics_agentcopilots_aggregates_jobsCmd.AddCommand(analytics_agentcopilots_aggregates_jobs_results.Cmdanalytics_agentcopilots_aggregates_jobs_results())
	analytics_agentcopilots_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_agentcopilots_aggregates_jobsCmd.Short, analytics_agentcopilots_aggregates_jobs_results.Description, )
	analytics_agentcopilots_aggregates_jobsCmd.Long = analytics_agentcopilots_aggregates_jobsCmd.Short
}
