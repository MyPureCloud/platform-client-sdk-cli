package analytics_agenticvirtualagents_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_agenticvirtualagents_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/agenticvirtualagents/aggregates")
	analytics_agenticvirtualagents_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_agenticvirtualagents_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_agenticvirtualagents_aggregatesCmd)
}

func Cmdanalytics_agenticvirtualagents_aggregates() *cobra.Command {
	return analytics_agenticvirtualagents_aggregatesCmd
}
