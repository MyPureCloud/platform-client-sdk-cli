package analytics_taskmanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_taskmanagement_aggregates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_taskmanagement_metrics"
)

func init() {
	analytics_taskmanagementCmd.AddCommand(analytics_taskmanagement_aggregates.Cmdanalytics_taskmanagement_aggregates())
	analytics_taskmanagementCmd.AddCommand(analytics_taskmanagement_metrics.Cmdanalytics_taskmanagement_metrics())
	analytics_taskmanagementCmd.Short = utils.GenerateCustomDescription(analytics_taskmanagementCmd.Short, analytics_taskmanagement_aggregates.Description, analytics_taskmanagement_metrics.Description, )
	analytics_taskmanagementCmd.Long = analytics_taskmanagementCmd.Short
}
