package analytics_taskmanagement_metrics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_taskmanagement_metrics_query"
)

func init() {
	analytics_taskmanagement_metricsCmd.AddCommand(analytics_taskmanagement_metrics_query.Cmdanalytics_taskmanagement_metrics_query())
	analytics_taskmanagement_metricsCmd.Short = utils.GenerateCustomDescription(analytics_taskmanagement_metricsCmd.Short, analytics_taskmanagement_metrics_query.Description, )
	analytics_taskmanagement_metricsCmd.Long = analytics_taskmanagement_metricsCmd.Short
}
