package analytics_surveys_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_surveys_aggregates_jobs_results"
)

func init() {
	analytics_surveys_aggregates_jobsCmd.AddCommand(analytics_surveys_aggregates_jobs_results.Cmdanalytics_surveys_aggregates_jobs_results())
	analytics_surveys_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_surveys_aggregates_jobsCmd.Short, analytics_surveys_aggregates_jobs_results.Description, )
	analytics_surveys_aggregates_jobsCmd.Long = analytics_surveys_aggregates_jobsCmd.Short
}
