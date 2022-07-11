package gamification_scorecards_profiles_metrics_values

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("gamification_scorecards_profiles_metrics_values", "SWAGGER_OVERRIDE_/api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/values")
	gamification_scorecards_profiles_metrics_valuesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("gamification_scorecards_profiles_metrics_values"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(gamification_scorecards_profiles_metrics_valuesCmd)
}

func Cmdgamification_scorecards_profiles_metrics_values() *cobra.Command {
	return gamification_scorecards_profiles_metrics_valuesCmd
}
