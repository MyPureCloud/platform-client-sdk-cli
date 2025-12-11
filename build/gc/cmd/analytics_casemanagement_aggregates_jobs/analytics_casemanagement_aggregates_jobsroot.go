package analytics_casemanagement_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_casemanagement_aggregates_jobs_results"
)

func init() {
	analytics_casemanagement_aggregates_jobsCmd.AddCommand(analytics_casemanagement_aggregates_jobs_results.Cmdanalytics_casemanagement_aggregates_jobs_results())
	analytics_casemanagement_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_casemanagement_aggregates_jobsCmd.Short, analytics_casemanagement_aggregates_jobs_results.Description, )
	analytics_casemanagement_aggregates_jobsCmd.Long = analytics_casemanagement_aggregates_jobsCmd.Short
}
