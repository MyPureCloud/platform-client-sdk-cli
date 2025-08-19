package analytics_knowledge_aggregates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_knowledge_aggregates_jobs_results"
)

func init() {
	analytics_knowledge_aggregates_jobsCmd.AddCommand(analytics_knowledge_aggregates_jobs_results.Cmdanalytics_knowledge_aggregates_jobs_results())
	analytics_knowledge_aggregates_jobsCmd.Short = utils.GenerateCustomDescription(analytics_knowledge_aggregates_jobsCmd.Short, analytics_knowledge_aggregates_jobs_results.Description, )
	analytics_knowledge_aggregates_jobsCmd.Long = analytics_knowledge_aggregates_jobsCmd.Short
}
