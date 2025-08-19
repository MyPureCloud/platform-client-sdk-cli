package analytics_reporting

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_reporting", "SWAGGER_OVERRIDE_/api/v2/analytics/reporting")
	analytics_reportingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_reporting"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_reportingCmd)
}

func Cmdanalytics_reporting() *cobra.Command {
	return analytics_reportingCmd
}
