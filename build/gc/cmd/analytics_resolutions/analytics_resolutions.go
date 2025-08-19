package analytics_resolutions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_resolutions", "SWAGGER_OVERRIDE_/api/v2/analytics/resolutions")
	analytics_resolutionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_resolutions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_resolutionsCmd)
}

func Cmdanalytics_resolutions() *cobra.Command {
	return analytics_resolutionsCmd
}
