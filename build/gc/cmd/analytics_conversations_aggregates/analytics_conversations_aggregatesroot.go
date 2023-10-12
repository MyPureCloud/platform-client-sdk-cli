package analytics_conversations_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_conversations_aggregates_jobs"
)

func init() {
	analytics_conversations_aggregatesCmd.AddCommand(analytics_conversations_aggregates_query.Cmdanalytics_conversations_aggregates_query())
	analytics_conversations_aggregatesCmd.AddCommand(analytics_conversations_aggregates_jobs.Cmdanalytics_conversations_aggregates_jobs())
	analytics_conversations_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_conversations_aggregatesCmd.Short, analytics_conversations_aggregates_query.Description, analytics_conversations_aggregates_jobs.Description, )
	analytics_conversations_aggregatesCmd.Long = analytics_conversations_aggregatesCmd.Short
}
