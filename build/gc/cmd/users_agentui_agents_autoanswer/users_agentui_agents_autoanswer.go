package users_agentui_agents_autoanswer

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("users_agentui_agents_autoanswer", "SWAGGER_OVERRIDE_/api/v2/users/agentui/agents/autoanswer")
	users_agentui_agents_autoanswerCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("users_agentui_agents_autoanswer"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(users_agentui_agents_autoanswerCmd)
}

func Cmdusers_agentui_agents_autoanswer() *cobra.Command {
	return users_agentui_agents_autoanswerCmd
}
