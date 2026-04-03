package analytics_copilots_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_copilots_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/copilots/aggregates")
	analytics_copilots_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_copilots_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_copilots_aggregatesCmd)
}

func Cmdanalytics_copilots_aggregates() *cobra.Command {
	return analytics_copilots_aggregatesCmd
}
