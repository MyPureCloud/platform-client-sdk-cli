package analytics_reporting_dashboards

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_reporting_dashboards", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting/dashboards")
	analytics_reporting_dashboardsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting_dashboards"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reporting_dashboardsCmd)
}

func Cmdanalytics_reporting_dashboards() *cobra.Command {
	return analytics_reporting_dashboardsCmd
}
