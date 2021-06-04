package analytics_journeys_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_journeys_aggregates_query"
)

func init() {
	analytics_journeys_aggregatesCmd.AddCommand(analytics_journeys_aggregates_query.Cmdanalytics_journeys_aggregates_query())
	analytics_journeys_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_journeys_aggregatesCmd.Short, analytics_journeys_aggregates_query.Description, )
	analytics_journeys_aggregatesCmd.Long = analytics_journeys_aggregatesCmd.Short
}
