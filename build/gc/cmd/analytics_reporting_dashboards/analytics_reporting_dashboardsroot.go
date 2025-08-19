package analytics_reporting_dashboards

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_dashboards_users"
)

func init() {
	analytics_reporting_dashboardsCmd.AddCommand(analytics_reporting_dashboards_users.Cmdanalytics_reporting_dashboards_users())
	analytics_reporting_dashboardsCmd.Short = utils.GenerateCustomDescription(analytics_reporting_dashboardsCmd.Short, analytics_reporting_dashboards_users.Description, )
	analytics_reporting_dashboardsCmd.Long = analytics_reporting_dashboardsCmd.Short
}
