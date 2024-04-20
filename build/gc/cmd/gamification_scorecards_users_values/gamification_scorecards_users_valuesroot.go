package gamification_scorecards_users_values

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_values_average"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_values_trends"
)

func init() {
	gamification_scorecards_users_valuesCmd.AddCommand(gamification_scorecards_users_values_average.Cmdgamification_scorecards_users_values_average())
	gamification_scorecards_users_valuesCmd.AddCommand(gamification_scorecards_users_values_trends.Cmdgamification_scorecards_users_values_trends())
	gamification_scorecards_users_valuesCmd.Short = utils.GenerateCustomDescription(gamification_scorecards_users_valuesCmd.Short, gamification_scorecards_users_values_average.Description, gamification_scorecards_users_values_trends.Description, )
	gamification_scorecards_users_valuesCmd.Long = gamification_scorecards_users_valuesCmd.Short
}
