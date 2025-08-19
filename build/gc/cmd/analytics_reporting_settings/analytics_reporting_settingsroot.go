package analytics_reporting_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_settings_dashboards"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/analytics_reporting_settings_users"
)

func init() {
	analytics_reporting_settingsCmd.AddCommand(analytics_reporting_settings_dashboards.Cmdanalytics_reporting_settings_dashboards())
	analytics_reporting_settingsCmd.AddCommand(analytics_reporting_settings_users.Cmdanalytics_reporting_settings_users())
	analytics_reporting_settingsCmd.Short = utils.GenerateCustomDescription(analytics_reporting_settingsCmd.Short, analytics_reporting_settings_dashboards.Description, analytics_reporting_settings_users.Description, )
	analytics_reporting_settingsCmd.Long = analytics_reporting_settingsCmd.Short
}
