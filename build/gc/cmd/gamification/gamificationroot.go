package gamification

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_status"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_leaderboard"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_metricdefinitions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_templates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards"
)

func init() {
	gamificationCmd.AddCommand(gamification_status.Cmdgamification_status())
	gamificationCmd.AddCommand(gamification_leaderboard.Cmdgamification_leaderboard())
	gamificationCmd.AddCommand(gamification_metricdefinitions.Cmdgamification_metricdefinitions())
	gamificationCmd.AddCommand(gamification_templates.Cmdgamification_templates())
	gamificationCmd.AddCommand(gamification_profiles.Cmdgamification_profiles())
	gamificationCmd.AddCommand(gamification_scorecards.Cmdgamification_scorecards())
	gamificationCmd.Short = utils.GenerateCustomDescription(gamificationCmd.Short, gamification_status.Description, gamification_leaderboard.Description, gamification_metricdefinitions.Description, gamification_templates.Description, gamification_profiles.Description, gamification_scorecards.Description, )
	gamificationCmd.Long = gamificationCmd.Short
}
