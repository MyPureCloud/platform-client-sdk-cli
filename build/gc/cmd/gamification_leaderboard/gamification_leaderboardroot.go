package gamification_leaderboard

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_leaderboard_all"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_leaderboard_bestpoints"
)

func init() {
	gamification_leaderboardCmd.AddCommand(gamification_leaderboard_all.Cmdgamification_leaderboard_all())
	gamification_leaderboardCmd.AddCommand(gamification_leaderboard_bestpoints.Cmdgamification_leaderboard_bestpoints())
	gamification_leaderboardCmd.Short = utils.GenerateCustomDescription(gamification_leaderboardCmd.Short, gamification_leaderboard_all.Description, gamification_leaderboard_bestpoints.Description, )
	gamification_leaderboardCmd.Long = gamification_leaderboardCmd.Short
}
