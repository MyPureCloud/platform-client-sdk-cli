package gamification_insights_groups_trends

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_groups_trends_all"
)

func init() {
	gamification_insights_groups_trendsCmd.AddCommand(gamification_insights_groups_trends_all.Cmdgamification_insights_groups_trends_all())
	gamification_insights_groups_trendsCmd.Short = utils.GenerateCustomDescription(gamification_insights_groups_trendsCmd.Short, gamification_insights_groups_trends_all.Description, )
	gamification_insights_groups_trendsCmd.Long = gamification_insights_groups_trendsCmd.Short
}
