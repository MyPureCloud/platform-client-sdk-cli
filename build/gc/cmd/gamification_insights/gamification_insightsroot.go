package gamification_insights

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_details"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_members"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_managers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_groups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_rankings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_trends"
)

func init() {
	gamification_insightsCmd.AddCommand(gamification_insights_users.Cmdgamification_insights_users())
	gamification_insightsCmd.AddCommand(gamification_insights_details.Cmdgamification_insights_details())
	gamification_insightsCmd.AddCommand(gamification_insights_members.Cmdgamification_insights_members())
	gamification_insightsCmd.AddCommand(gamification_insights_managers.Cmdgamification_insights_managers())
	gamification_insightsCmd.AddCommand(gamification_insights_groups.Cmdgamification_insights_groups())
	gamification_insightsCmd.AddCommand(gamification_insights_rankings.Cmdgamification_insights_rankings())
	gamification_insightsCmd.AddCommand(gamification_insights_trends.Cmdgamification_insights_trends())
	gamification_insightsCmd.Short = utils.GenerateCustomDescription(gamification_insightsCmd.Short, gamification_insights_users.Description, gamification_insights_details.Description, gamification_insights_members.Description, gamification_insights_managers.Description, gamification_insights_groups.Description, gamification_insights_rankings.Description, gamification_insights_trends.Description, )
	gamification_insightsCmd.Long = gamification_insightsCmd.Short
}
