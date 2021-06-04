package analytics_reporting_schedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_schedules_history"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_schedules_runreport"
)

func init() {
	analytics_reporting_schedulesCmd.AddCommand(analytics_reporting_schedules_history.Cmdanalytics_reporting_schedules_history())
	analytics_reporting_schedulesCmd.AddCommand(analytics_reporting_schedules_runreport.Cmdanalytics_reporting_schedules_runreport())
	analytics_reporting_schedulesCmd.Short = utils.GenerateCustomDescription(analytics_reporting_schedulesCmd.Short, analytics_reporting_schedules_history.Description, analytics_reporting_schedules_runreport.Description, )
	analytics_reporting_schedulesCmd.Long = analytics_reporting_schedulesCmd.Short
}
