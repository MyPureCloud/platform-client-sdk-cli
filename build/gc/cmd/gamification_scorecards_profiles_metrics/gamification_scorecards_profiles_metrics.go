package gamification_scorecards_profiles_metrics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_scorecards_profiles_metrics", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/profiles/{profileId}/metrics")
	gamification_scorecards_profiles_metricsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_profiles_metrics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_profiles_metricsCmd)
}

func Cmdgamification_scorecards_profiles_metrics() *cobra.Command {
	return gamification_scorecards_profiles_metricsCmd
}
