package gamification_contests_agents_scores_trends

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_agents_scores_trends_me"
)

func init() {
	gamification_contests_agents_scores_trendsCmd.AddCommand(gamification_contests_agents_scores_trends_me.Cmdgamification_contests_agents_scores_trends_me())
	gamification_contests_agents_scores_trendsCmd.Short = utils.GenerateCustomDescription(gamification_contests_agents_scores_trendsCmd.Short, gamification_contests_agents_scores_trends_me.Description, )
	gamification_contests_agents_scores_trendsCmd.Long = gamification_contests_agents_scores_trendsCmd.Short
}
