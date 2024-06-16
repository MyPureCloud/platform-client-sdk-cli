package analytics_agentcopilots_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agentcopilots_aggregates_jobs"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_agentcopilots_aggregates_query"
)

func init() {
	analytics_agentcopilots_aggregatesCmd.AddCommand(analytics_agentcopilots_aggregates_jobs.Cmdanalytics_agentcopilots_aggregates_jobs())
	analytics_agentcopilots_aggregatesCmd.AddCommand(analytics_agentcopilots_aggregates_query.Cmdanalytics_agentcopilots_aggregates_query())
	analytics_agentcopilots_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_agentcopilots_aggregatesCmd.Short, analytics_agentcopilots_aggregates_jobs.Description, analytics_agentcopilots_aggregates_query.Description, )
	analytics_agentcopilots_aggregatesCmd.Long = analytics_agentcopilots_aggregatesCmd.Short
}
