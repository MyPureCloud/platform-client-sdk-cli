package analytics_reporting

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_dashboards"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_exports"
)

func init() {
	analytics_reportingCmd.AddCommand(analytics_reporting_settings.Cmdanalytics_reporting_settings())
	analytics_reportingCmd.AddCommand(analytics_reporting_dashboards.Cmdanalytics_reporting_dashboards())
	analytics_reportingCmd.AddCommand(analytics_reporting_exports.Cmdanalytics_reporting_exports())
	analytics_reportingCmd.Short = utils.GenerateCustomDescription(analytics_reportingCmd.Short, analytics_reporting_settings.Description, analytics_reporting_dashboards.Description, analytics_reporting_exports.Description, )
	analytics_reportingCmd.Long = analytics_reportingCmd.Short
}
