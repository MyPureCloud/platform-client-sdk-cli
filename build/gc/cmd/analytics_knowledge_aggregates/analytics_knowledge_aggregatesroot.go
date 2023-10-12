package analytics_knowledge_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_knowledge_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_knowledge_aggregates_jobs"
)

func init() {
	analytics_knowledge_aggregatesCmd.AddCommand(analytics_knowledge_aggregates_query.Cmdanalytics_knowledge_aggregates_query())
	analytics_knowledge_aggregatesCmd.AddCommand(analytics_knowledge_aggregates_jobs.Cmdanalytics_knowledge_aggregates_jobs())
	analytics_knowledge_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_knowledge_aggregatesCmd.Short, analytics_knowledge_aggregates_query.Description, analytics_knowledge_aggregates_jobs.Description, )
	analytics_knowledge_aggregatesCmd.Long = analytics_knowledge_aggregatesCmd.Short
}
