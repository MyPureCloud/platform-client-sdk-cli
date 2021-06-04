package analytics_journeys

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_journeys_aggregates"
)

func init() {
	analytics_journeysCmd.AddCommand(analytics_journeys_aggregates.Cmdanalytics_journeys_aggregates())
	analytics_journeysCmd.Short = utils.GenerateCustomDescription(analytics_journeysCmd.Short, analytics_journeys_aggregates.Description, )
	analytics_journeysCmd.Long = analytics_journeysCmd.Short
}
