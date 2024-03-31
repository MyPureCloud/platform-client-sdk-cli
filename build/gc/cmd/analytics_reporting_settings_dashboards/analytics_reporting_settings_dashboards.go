package analytics_reporting_settings_dashboards

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_reporting_settings_dashboards", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting/settings/dashboards")
	analytics_reporting_settings_dashboardsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting_settings_dashboards"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reporting_settings_dashboardsCmd)
}

func Cmdanalytics_reporting_settings_dashboards() *cobra.Command {
	return analytics_reporting_settings_dashboardsCmd
}
