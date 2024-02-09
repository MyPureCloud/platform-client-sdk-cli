package analytics_ratelimits_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_ratelimits_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/ratelimits/aggregates")
	analytics_ratelimits_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_ratelimits_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_ratelimits_aggregatesCmd)
}

func Cmdanalytics_ratelimits_aggregates() *cobra.Command {
	return analytics_ratelimits_aggregatesCmd
}
