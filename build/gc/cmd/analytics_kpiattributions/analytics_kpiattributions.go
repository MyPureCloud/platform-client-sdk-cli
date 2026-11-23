package analytics_kpiattributions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_kpiattributions", "SWAGGER_OVERRIDE_/api/v2/analytics/kpiattributions")
	analytics_kpiattributionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_kpiattributions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_kpiattributionsCmd)
}

func Cmdanalytics_kpiattributions() *cobra.Command {
	return analytics_kpiattributionsCmd
}
