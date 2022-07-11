package gamification_scorecards_profiles_metrics_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_scorecards_profiles_metrics_users", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users")
	gamification_scorecards_profiles_metrics_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_profiles_metrics_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_profiles_metrics_usersCmd)
}

func Cmdgamification_scorecards_profiles_metrics_users() *cobra.Command {
	return gamification_scorecards_profiles_metrics_usersCmd
}
