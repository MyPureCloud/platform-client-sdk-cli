package analytics_agentcopilots_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_agentcopilots_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/agentcopilots/aggregates")
	analytics_agentcopilots_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_agentcopilots_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_agentcopilots_aggregatesCmd)
}

func Cmdanalytics_agentcopilots_aggregates() *cobra.Command {
	return analytics_agentcopilots_aggregatesCmd
}
