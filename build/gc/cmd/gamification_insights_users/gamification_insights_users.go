package gamification_insights_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_insights_users", "SWAGGER_OVERRIDE_/api/v2/gamification/insights/users")
	gamification_insights_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_insights_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_insights_usersCmd)
}

func Cmdgamification_insights_users() *cobra.Command {
	return gamification_insights_usersCmd
}
