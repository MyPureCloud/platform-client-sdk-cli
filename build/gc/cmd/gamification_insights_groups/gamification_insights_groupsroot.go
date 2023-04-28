package gamification_insights_groups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_groups_trends"
)

func init() {
	gamification_insights_groupsCmd.AddCommand(gamification_insights_groups_trends.Cmdgamification_insights_groups_trends())
	gamification_insights_groupsCmd.Short = utils.GenerateCustomDescription(gamification_insights_groupsCmd.Short, gamification_insights_groups_trends.Description, )
	gamification_insights_groupsCmd.Long = gamification_insights_groupsCmd.Short
}
