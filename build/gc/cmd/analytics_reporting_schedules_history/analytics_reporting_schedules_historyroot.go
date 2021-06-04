package analytics_reporting_schedules_history

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_schedules_history_latest"
)

func init() {
	analytics_reporting_schedules_historyCmd.AddCommand(analytics_reporting_schedules_history_latest.Cmdanalytics_reporting_schedules_history_latest())
	analytics_reporting_schedules_historyCmd.Short = utils.GenerateCustomDescription(analytics_reporting_schedules_historyCmd.Short, analytics_reporting_schedules_history_latest.Description, )
	analytics_reporting_schedules_historyCmd.Long = analytics_reporting_schedules_historyCmd.Short
}
