package analytics_actions_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_actions_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_actions_aggregates_jobs"
)

func init() {
	analytics_actions_aggregatesCmd.AddCommand(analytics_actions_aggregates_query.Cmdanalytics_actions_aggregates_query())
	analytics_actions_aggregatesCmd.AddCommand(analytics_actions_aggregates_jobs.Cmdanalytics_actions_aggregates_jobs())
	analytics_actions_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_actions_aggregatesCmd.Short, analytics_actions_aggregates_query.Description, analytics_actions_aggregates_jobs.Description, )
	analytics_actions_aggregatesCmd.Long = analytics_actions_aggregatesCmd.Short
}
