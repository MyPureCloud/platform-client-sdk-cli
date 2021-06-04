package gamification_leaderboard_all

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_leaderboard_all_bestpoints"
)

func init() {
	gamification_leaderboard_allCmd.AddCommand(gamification_leaderboard_all_bestpoints.Cmdgamification_leaderboard_all_bestpoints())
	gamification_leaderboard_allCmd.Short = utils.GenerateCustomDescription(gamification_leaderboard_allCmd.Short, gamification_leaderboard_all_bestpoints.Description, )
	gamification_leaderboard_allCmd.Long = gamification_leaderboard_allCmd.Short
}
