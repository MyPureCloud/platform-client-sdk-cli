package gamification_profiles

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_activate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_deactivate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_metrics"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_members"
)

func init() {
	gamification_profilesCmd.AddCommand(gamification_profiles_activate.Cmdgamification_profiles_activate())
	gamification_profilesCmd.AddCommand(gamification_profiles_deactivate.Cmdgamification_profiles_deactivate())
	gamification_profilesCmd.AddCommand(gamification_profiles_users.Cmdgamification_profiles_users())
	gamification_profilesCmd.AddCommand(gamification_profiles_metrics.Cmdgamification_profiles_metrics())
	gamification_profilesCmd.AddCommand(gamification_profiles_members.Cmdgamification_profiles_members())
	gamification_profilesCmd.Short = utils.GenerateCustomDescription(gamification_profilesCmd.Short, gamification_profiles_activate.Description, gamification_profiles_deactivate.Description, gamification_profiles_users.Description, gamification_profiles_metrics.Description, gamification_profiles_members.Description, )
	gamification_profilesCmd.Long = gamification_profilesCmd.Short
}
