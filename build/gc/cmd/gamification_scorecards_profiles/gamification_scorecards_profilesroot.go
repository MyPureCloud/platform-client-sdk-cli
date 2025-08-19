package gamification_scorecards_profiles

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_profiles_metrics"
)

func init() {
	gamification_scorecards_profilesCmd.AddCommand(gamification_scorecards_profiles_metrics.Cmdgamification_scorecards_profiles_metrics())
	gamification_scorecards_profilesCmd.Short = utils.GenerateCustomDescription(gamification_scorecards_profilesCmd.Short, gamification_scorecards_profiles_metrics.Description, )
	gamification_scorecards_profilesCmd.Long = gamification_scorecards_profilesCmd.Short
}
