package gamification_scorecards_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_bestpoints"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_attendance"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_values"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users_points"
)

func init() {
	gamification_scorecards_usersCmd.AddCommand(gamification_scorecards_users_bestpoints.Cmdgamification_scorecards_users_bestpoints())
	gamification_scorecards_usersCmd.AddCommand(gamification_scorecards_users_attendance.Cmdgamification_scorecards_users_attendance())
	gamification_scorecards_usersCmd.AddCommand(gamification_scorecards_users_values.Cmdgamification_scorecards_users_values())
	gamification_scorecards_usersCmd.AddCommand(gamification_scorecards_users_points.Cmdgamification_scorecards_users_points())
	gamification_scorecards_usersCmd.Short = utils.GenerateCustomDescription(gamification_scorecards_usersCmd.Short, gamification_scorecards_users_bestpoints.Description, gamification_scorecards_users_attendance.Description, gamification_scorecards_users_values.Description, gamification_scorecards_users_points.Description, )
	gamification_scorecards_usersCmd.Long = gamification_scorecards_usersCmd.Short
}
