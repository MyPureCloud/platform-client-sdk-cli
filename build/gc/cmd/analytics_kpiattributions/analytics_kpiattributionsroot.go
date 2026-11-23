package analytics_kpiattributions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_kpiattributions_aggregates"
)

func init() {
	analytics_kpiattributionsCmd.AddCommand(analytics_kpiattributions_aggregates.Cmdanalytics_kpiattributions_aggregates())
	analytics_kpiattributionsCmd.Short = utils.GenerateCustomDescription(analytics_kpiattributionsCmd.Short, analytics_kpiattributions_aggregates.Description, )
	analytics_kpiattributionsCmd.Long = analytics_kpiattributionsCmd.Short
}
