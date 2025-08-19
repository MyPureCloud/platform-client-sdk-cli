package gamification_profiles_members

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/gamification_profiles_members_validate"
)

func init() {
	gamification_profiles_membersCmd.AddCommand(gamification_profiles_members_validate.Cmdgamification_profiles_members_validate())
	gamification_profiles_membersCmd.Short = utils.GenerateCustomDescription(gamification_profiles_membersCmd.Short, gamification_profiles_members_validate.Description, )
	gamification_profiles_membersCmd.Long = gamification_profiles_membersCmd.Short
}
