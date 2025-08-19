package analytics_resolutions_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_resolutions_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/resolutions/aggregates")
	analytics_resolutions_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_resolutions_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_resolutions_aggregatesCmd)
}

func Cmdanalytics_resolutions_aggregates() *cobra.Command {
	return analytics_resolutions_aggregatesCmd
}
