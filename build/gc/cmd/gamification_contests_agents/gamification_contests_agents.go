package gamification_contests_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_contests_agents", "SWAGGER_OVERRIDE_/api/v2/gamification/contests/{contestId}/agents")
	gamification_contests_agentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_contests_agents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_contests_agentsCmd)
}

func Cmdgamification_contests_agents() *cobra.Command {
	return gamification_contests_agentsCmd
}
