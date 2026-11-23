package analytics_kpiattributions_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_kpiattributions_aggregates_jobs_results"
)

func init() {
	analytics_kpiattributions_aggregates_jobsCmd.AddCommand(analytics_kpiattributions_aggregates_jobs_results.Cmdanalytics_kpiattributions_aggregates_jobs_results())
	analytics_kpiattributions_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_kpiattributions_aggregates_jobsCmd.Short, analytics_kpiattributions_aggregates_jobs_results.Description, )
	analytics_kpiattributions_aggregates_jobsCmd.Long = analytics_kpiattributions_aggregates_jobsCmd.Short
}
