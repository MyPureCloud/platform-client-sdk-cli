package analytics_summaries_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_summaries_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_summaries_aggregates_jobs"
)

func init() {
	analytics_summaries_aggregatesCmd.AddCommand(analytics_summaries_aggregates_query.Cmdanalytics_summaries_aggregates_query())
	analytics_summaries_aggregatesCmd.AddCommand(analytics_summaries_aggregates_jobs.Cmdanalytics_summaries_aggregates_jobs())
	analytics_summaries_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_summaries_aggregatesCmd.Short, analytics_summaries_aggregates_query.Description, analytics_summaries_aggregates_jobs.Description, )
	analytics_summaries_aggregatesCmd.Long = analytics_summaries_aggregatesCmd.Short
}
