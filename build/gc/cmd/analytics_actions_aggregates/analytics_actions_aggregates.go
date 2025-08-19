package analytics_actions_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_actions_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/actions/aggregates")
	analytics_actions_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_actions_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_actions_aggregatesCmd)
}

func Cmdanalytics_actions_aggregates() *cobra.Command {
	return analytics_actions_aggregatesCmd
}
