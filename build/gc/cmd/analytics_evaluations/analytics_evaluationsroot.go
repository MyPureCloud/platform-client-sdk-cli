package analytics_evaluations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_evaluations_aggregates"
)

func init() {
	analytics_evaluationsCmd.AddCommand(analytics_evaluations_aggregates.Cmdanalytics_evaluations_aggregates())
	analytics_evaluationsCmd.Short = utils.GenerateCustomDescription(analytics_evaluationsCmd.Short, analytics_evaluations_aggregates.Description, )
	analytics_evaluationsCmd.Long = analytics_evaluationsCmd.Short
}
