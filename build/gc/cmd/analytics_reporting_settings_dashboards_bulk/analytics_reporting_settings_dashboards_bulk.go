package analytics_reporting_settings_dashboards_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_reporting_settings_dashboards_bulk", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting/settings/dashboards/bulk")
	analytics_reporting_settings_dashboards_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting_settings_dashboards_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reporting_settings_dashboards_bulkCmd)
}

func Cmdanalytics_reporting_settings_dashboards_bulk() *cobra.Command {
	return analytics_reporting_settings_dashboards_bulkCmd
}
