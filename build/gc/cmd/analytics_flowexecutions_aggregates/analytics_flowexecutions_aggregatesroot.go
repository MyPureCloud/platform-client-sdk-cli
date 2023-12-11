package analytics_flowexecutions_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flowexecutions_aggregates_query"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flowexecutions_aggregates_jobs"
)

func init() {
	analytics_flowexecutions_aggregatesCmd.AddCommand(analytics_flowexecutions_aggregates_query.Cmdanalytics_flowexecutions_aggregates_query())
	analytics_flowexecutions_aggregatesCmd.AddCommand(analytics_flowexecutions_aggregates_jobs.Cmdanalytics_flowexecutions_aggregates_jobs())
	analytics_flowexecutions_aggregatesCmd.Short = utils.GenerateCustomDescription(analytics_flowexecutions_aggregatesCmd.Short, analytics_flowexecutions_aggregates_query.Description, analytics_flowexecutions_aggregates_jobs.Description, )
	analytics_flowexecutions_aggregatesCmd.Long = analytics_flowexecutions_aggregatesCmd.Short
}
