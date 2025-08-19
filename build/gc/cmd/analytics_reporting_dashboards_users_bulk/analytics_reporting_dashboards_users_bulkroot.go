package analytics_reporting_dashboards_users_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_dashboards_users_bulk_remove"
)

func init() {
	analytics_reporting_dashboards_users_bulkCmd.AddCommand(analytics_reporting_dashboards_users_bulk_remove.Cmdanalytics_reporting_dashboards_users_bulk_remove())
	analytics_reporting_dashboards_users_bulkCmd.Short = utils.GenerateCustomDescription(analytics_reporting_dashboards_users_bulkCmd.Short, analytics_reporting_dashboards_users_bulk_remove.Description, )
	analytics_reporting_dashboards_users_bulkCmd.Long = analytics_reporting_dashboards_users_bulkCmd.Short
}
