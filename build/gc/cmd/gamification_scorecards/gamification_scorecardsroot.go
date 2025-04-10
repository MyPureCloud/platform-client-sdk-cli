package gamification_scorecards

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_profiles"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_points"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_values"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_attendance"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_bestpoints"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_scorecards_users"
)

func init() {
	gamification_scorecardsCmd.AddCommand(gamification_scorecards_profiles.Cmdgamification_scorecards_profiles())
	gamification_scorecardsCmd.AddCommand(gamification_scorecards_points.Cmdgamification_scorecards_points())
	gamification_scorecardsCmd.AddCommand(gamification_scorecards_values.Cmdgamification_scorecards_values())
	gamification_scorecardsCmd.AddCommand(gamification_scorecards_attendance.Cmdgamification_scorecards_attendance())
	gamification_scorecardsCmd.AddCommand(gamification_scorecards_bestpoints.Cmdgamification_scorecards_bestpoints())
	gamification_scorecardsCmd.AddCommand(gamification_scorecards_users.Cmdgamification_scorecards_users())
	gamification_scorecardsCmd.Short = utils.GenerateCustomDescription(gamification_scorecardsCmd.Short, gamification_scorecards_profiles.Description, gamification_scorecards_points.Description, gamification_scorecards_values.Description, gamification_scorecards_attendance.Description, gamification_scorecards_bestpoints.Description, gamification_scorecards_users.Description, )
	gamification_scorecardsCmd.Long = gamification_scorecardsCmd.Short
}
