package analytics_reporting_dashboards_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_dashboards_users_bulk"
)

func init() {
	analytics_reporting_dashboards_usersCmd.AddCommand(analytics_reporting_dashboards_users_bulk.Cmdanalytics_reporting_dashboards_users_bulk())
	analytics_reporting_dashboards_usersCmd.Short = utils.GenerateCustomDescription(analytics_reporting_dashboards_usersCmd.Short, analytics_reporting_dashboards_users_bulk.Description, )
	analytics_reporting_dashboards_usersCmd.Long = analytics_reporting_dashboards_usersCmd.Short
}
