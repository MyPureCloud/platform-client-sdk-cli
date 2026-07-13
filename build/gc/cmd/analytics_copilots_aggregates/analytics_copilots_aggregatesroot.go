package analytics_copilots_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_copilots_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_copilots_aggregates_jobs"
)

func init() {
	analytics_copilots_aggregatesCmd.AddCommand(analytics_copilots_aggregates_query.Cmdanalytics_copilots_aggregates_query())
	analytics_copilots_aggregatesCmd.AddCommand(analytics_copilots_aggregates_jobs.Cmdanalytics_copilots_aggregates_jobs())
	analytics_copilots_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_copilots_aggregatesCmd.Short, analytics_copilots_aggregates_query.Description, analytics_copilots_aggregates_jobs.Description, )
	analytics_copilots_aggregatesCmd.Long = analytics_copilots_aggregatesCmd.Short
}
