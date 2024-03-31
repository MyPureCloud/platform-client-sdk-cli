package analytics_reporting_settings_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_settings_users_dashboards"
)

func init() {
	analytics_reporting_settings_usersCmd.AddCommand(analytics_reporting_settings_users_dashboards.Cmdanalytics_reporting_settings_users_dashboards())
	analytics_reporting_settings_usersCmd.Short = utils.GenerateCustomDescription(analytics_reporting_settings_usersCmd.Short, analytics_reporting_settings_users_dashboards.Description, )
	analytics_reporting_settings_usersCmd.Long = analytics_reporting_settings_usersCmd.Short
}
