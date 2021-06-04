package analytics_bots_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_bots_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/bots/aggregates")
	analytics_bots_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_bots_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_bots_aggregatesCmd)
}

func Cmdanalytics_bots_aggregates() *cobra.Command {
	return analytics_bots_aggregatesCmd
}
