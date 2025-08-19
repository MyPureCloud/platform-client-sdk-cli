package quality_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("quality_agents", "SWAGGER_OVERRIDE_/api/v2/quality/agents")
	quality_agentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("quality_agents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(quality_agentsCmd)
}

func Cmdquality_agents() *cobra.Command {
	return quality_agentsCmd
}
