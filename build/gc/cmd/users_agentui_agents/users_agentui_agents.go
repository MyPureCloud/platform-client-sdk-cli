package users_agentui_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_agentui_agents", "SWAGGER_OVERRIDE_/api/v2/users/agentui/agents")
	users_agentui_agentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_agentui_agents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_agentui_agentsCmd)
}

func Cmdusers_agentui_agents() *cobra.Command {
	return users_agentui_agentsCmd
}
