package analytics_reporting

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_exports"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_reportformats"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_schedules"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_metadata"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_timeperiods"
)

func init() {
	analytics_reportingCmd.AddCommand(analytics_reporting_exports.Cmdanalytics_reporting_exports())
	analytics_reportingCmd.AddCommand(analytics_reporting_reportformats.Cmdanalytics_reporting_reportformats())
	analytics_reportingCmd.AddCommand(analytics_reporting_schedules.Cmdanalytics_reporting_schedules())
	analytics_reportingCmd.AddCommand(analytics_reporting_metadata.Cmdanalytics_reporting_metadata())
	analytics_reportingCmd.AddCommand(analytics_reporting_timeperiods.Cmdanalytics_reporting_timeperiods())
	analytics_reportingCmd.Short = utils.GenerateCustomDescription(analytics_reportingCmd.Short, analytics_reporting_exports.Description, analytics_reporting_reportformats.Description, analytics_reporting_schedules.Description, analytics_reporting_metadata.Description, analytics_reporting_timeperiods.Description, )
	analytics_reportingCmd.Long = analytics_reportingCmd.Short
}
