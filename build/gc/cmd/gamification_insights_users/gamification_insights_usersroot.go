package gamification_insights_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_users_details"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_insights_users_trends"
)

func init() {
	gamification_insights_usersCmd.AddCommand(gamification_insights_users_details.Cmdgamification_insights_users_details())
	gamification_insights_usersCmd.AddCommand(gamification_insights_users_trends.Cmdgamification_insights_users_trends())
	gamification_insights_usersCmd.Short = utils.GenerateCustomDescription(gamification_insights_usersCmd.Short, gamification_insights_users_details.Description, gamification_insights_users_trends.Description, )
	gamification_insights_usersCmd.Long = gamification_insights_usersCmd.Short
}
