package analytics_flowexecutions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_flowexecutions_aggregates"
)

func init() {
	analytics_flowexecutionsCmd.AddCommand(analytics_flowexecutions_aggregates.Cmdanalytics_flowexecutions_aggregates())
	analytics_flowexecutionsCmd.Short = utils.GenerateCustomDescription(analytics_flowexecutionsCmd.Short, analytics_flowexecutions_aggregates.Description, )
	analytics_flowexecutionsCmd.Long = analytics_flowexecutionsCmd.Short
}
