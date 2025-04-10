package gamification_contests_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_agents_scores"
)

func init() {
	gamification_contests_agentsCmd.AddCommand(gamification_contests_agents_scores.Cmdgamification_contests_agents_scores())
	gamification_contests_agentsCmd.Short = utils.GenerateCustomDescription(gamification_contests_agentsCmd.Short, gamification_contests_agents_scores.Description, )
	gamification_contests_agentsCmd.Long = gamification_contests_agentsCmd.Short
}
