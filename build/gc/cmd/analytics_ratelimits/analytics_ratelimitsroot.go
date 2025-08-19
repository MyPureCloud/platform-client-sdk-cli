package analytics_ratelimits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_ratelimits_aggregates"
)

func init() {
	analytics_ratelimitsCmd.AddCommand(analytics_ratelimits_aggregates.Cmdanalytics_ratelimits_aggregates())
	analytics_ratelimitsCmd.Short = utils.GenerateCustomDescription(analytics_ratelimitsCmd.Short, analytics_ratelimits_aggregates.Description, )
	analytics_ratelimitsCmd.Long = analytics_ratelimitsCmd.Short
}
