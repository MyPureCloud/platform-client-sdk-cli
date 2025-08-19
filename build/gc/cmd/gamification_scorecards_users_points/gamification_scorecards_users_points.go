package gamification_scorecards_users_points

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_scorecards_users_points", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/users/{userId}/points")
	gamification_scorecards_users_pointsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_users_points"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_users_pointsCmd)
}

func Cmdgamification_scorecards_users_points() *cobra.Command {
	return gamification_scorecards_users_pointsCmd
}
