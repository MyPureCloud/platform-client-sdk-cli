package gamification_contests_agents_scores

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_agents_scores_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_agents_scores_trends"
)

func init() {
	gamification_contests_agents_scoresCmd.AddCommand(gamification_contests_agents_scores_me.Cmdgamification_contests_agents_scores_me())
	gamification_contests_agents_scoresCmd.AddCommand(gamification_contests_agents_scores_trends.Cmdgamification_contests_agents_scores_trends())
	gamification_contests_agents_scoresCmd.Short = utils.GenerateCustomDescription(gamification_contests_agents_scoresCmd.Short, gamification_contests_agents_scores_me.Description, gamification_contests_agents_scores_trends.Description, )
	gamification_contests_agents_scoresCmd.Long = gamification_contests_agents_scoresCmd.Short
}
