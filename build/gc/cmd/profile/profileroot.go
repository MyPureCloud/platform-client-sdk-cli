package profile

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/profile_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/profile_groups"
)

func init() {
	profileCmd.AddCommand(profile_users.Cmdprofile_users())
	profileCmd.AddCommand(profile_groups.Cmdprofile_groups())
	profileCmd.Short = utils.GenerateCustomDescription(profileCmd.Short, profile_users.Description, profile_groups.Description, )
	profileCmd.Long = profileCmd.Short
}
