package gamification_scorecards_profiles

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_scorecards_profiles", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/profiles")
	gamification_scorecards_profilesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_profiles"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_profilesCmd)
}

func Cmdgamification_scorecards_profiles() *cobra.Command {
	return gamification_scorecards_profilesCmd
}
