package analytics_ratelimits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_ratelimits", "SWAGGER_OVERRIDE_/api/v2/analytics/ratelimits")
	analytics_ratelimitsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_ratelimits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_ratelimitsCmd)
}

func Cmdanalytics_ratelimits() *cobra.Command {
	return analytics_ratelimitsCmd
}
