package gamification_profiles_users_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_users_me_query"
)

func init() {
	gamification_profiles_users_meCmd.AddCommand(gamification_profiles_users_me_query.Cmdgamification_profiles_users_me_query())
	gamification_profiles_users_meCmd.Short = utils.GenerateCustomDescription(gamification_profiles_users_meCmd.Short, gamification_profiles_users_me_query.Description, )
	gamification_profiles_users_meCmd.Long = gamification_profiles_users_meCmd.Short
}
