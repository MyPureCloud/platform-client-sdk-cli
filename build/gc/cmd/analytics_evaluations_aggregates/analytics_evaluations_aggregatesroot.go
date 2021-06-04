package analytics_evaluations_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_evaluations_aggregates_query"
)

func init() {
	analytics_evaluations_aggregatesCmd.AddCommand(analytics_evaluations_aggregates_query.Cmdanalytics_evaluations_aggregates_query())
	analytics_evaluations_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_evaluations_aggregatesCmd.Short, analytics_evaluations_aggregates_query.Description, )
	analytics_evaluations_aggregatesCmd.Long = analytics_evaluations_aggregatesCmd.Short
}
