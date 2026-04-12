package analytics_agentutilizations_aggregates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_agentutilizations_aggregates", "SWAGGER_OVERRIDE_/api/v2/analytics/agentutilizations/aggregates")
	analytics_agentutilizations_aggregatesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_agentutilizations_aggregates"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_agentutilizations_aggregatesCmd)
}

func Cmdanalytics_agentutilizations_aggregates() *cobra.Command {
	return analytics_agentutilizations_aggregatesCmd
}
