package analytics_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("analytics_agents", "SWAGGER_OVERRIDE_/api/v2/analytics/agents")
	analytics_agentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("analytics_agents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(analytics_agentsCmd)
}

func Cmdanalytics_agents() *cobra.Command {
	return analytics_agentsCmd
}
