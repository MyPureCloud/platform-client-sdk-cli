package gamification_scorecards_points

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_scorecards_points", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/points")
	gamification_scorecards_pointsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_points"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_pointsCmd)
}

func Cmdgamification_scorecards_points() *cobra.Command {
	return gamification_scorecards_pointsCmd
}
