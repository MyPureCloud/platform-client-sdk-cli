package analytics_kpiattributions_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_kpiattributions_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_kpiattributions_aggregates_jobs"
)

func init() {
	analytics_kpiattributions_aggregatesCmd.AddCommand(analytics_kpiattributions_aggregates_query.Cmdanalytics_kpiattributions_aggregates_query())
	analytics_kpiattributions_aggregatesCmd.AddCommand(analytics_kpiattributions_aggregates_jobs.Cmdanalytics_kpiattributions_aggregates_jobs())
	analytics_kpiattributions_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_kpiattributions_aggregatesCmd.Short, analytics_kpiattributions_aggregates_query.Description, analytics_kpiattributions_aggregates_jobs.Description, )
	analytics_kpiattributions_aggregatesCmd.Long = analytics_kpiattributions_aggregatesCmd.Short
}
