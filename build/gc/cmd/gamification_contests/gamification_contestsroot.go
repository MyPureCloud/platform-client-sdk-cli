package gamification_contests

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_uploads"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_prizeimages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_contests_agents"
)

func init() {
	gamification_contestsCmd.AddCommand(gamification_contests_uploads.Cmdgamification_contests_uploads())
	gamification_contestsCmd.AddCommand(gamification_contests_prizeimages.Cmdgamification_contests_prizeimages())
	gamification_contestsCmd.AddCommand(gamification_contests_me.Cmdgamification_contests_me())
	gamification_contestsCmd.AddCommand(gamification_contests_agents.Cmdgamification_contests_agents())
	gamification_contestsCmd.Short = utils.GenerateCustomDescription(gamification_contestsCmd.Short, gamification_contests_uploads.Description, gamification_contests_prizeimages.Description, gamification_contests_me.Description, gamification_contests_agents.Description, )
	gamification_contestsCmd.Long = gamification_contestsCmd.Short
}
