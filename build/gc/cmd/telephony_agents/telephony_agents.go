package telephony_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("telephony_agents", "SWAGGER_OVERRIDE_/api/v2/telephony/agents")
	telephony_agentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("telephony_agents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(telephony_agentsCmd)
}

func Cmdtelephony_agents() *cobra.Command {
	return telephony_agentsCmd
}
