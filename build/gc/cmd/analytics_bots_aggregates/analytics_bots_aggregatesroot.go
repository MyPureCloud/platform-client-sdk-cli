package analytics_bots_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_bots_aggregates_query"
)

func init() {
	analytics_bots_aggregatesCmd.AddCommand(analytics_bots_aggregates_query.Cmdanalytics_bots_aggregates_query())
	analytics_bots_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_bots_aggregatesCmd.Short, analytics_bots_aggregates_query.Description, )
	analytics_bots_aggregatesCmd.Long = analytics_bots_aggregatesCmd.Short
}
