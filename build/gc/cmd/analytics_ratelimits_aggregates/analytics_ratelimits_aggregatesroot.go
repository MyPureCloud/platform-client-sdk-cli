package analytics_ratelimits_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_ratelimits_aggregates_query"
)

func init() {
	analytics_ratelimits_aggregatesCmd.AddCommand(analytics_ratelimits_aggregates_query.Cmdanalytics_ratelimits_aggregates_query())
	analytics_ratelimits_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_ratelimits_aggregatesCmd.Short, analytics_ratelimits_aggregates_query.Description, )
	analytics_ratelimits_aggregatesCmd.Long = analytics_ratelimits_aggregatesCmd.Short
}
