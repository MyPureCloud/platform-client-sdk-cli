package analytics_flows_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flows_aggregates_jobs_results"
)

func init() {
	analytics_flows_aggregates_jobsCmd.AddCommand(analytics_flows_aggregates_jobs_results.Cmdanalytics_flows_aggregates_jobs_results())
	analytics_flows_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_flows_aggregates_jobsCmd.Short, analytics_flows_aggregates_jobs_results.Description, )
	analytics_flows_aggregates_jobsCmd.Long = analytics_flows_aggregates_jobsCmd.Short
}
