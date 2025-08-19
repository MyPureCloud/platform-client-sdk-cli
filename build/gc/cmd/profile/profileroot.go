package profile

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/profile_groups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/profile_users"
)

func init() {
	profileCmd.AddCommand(profile_groups.Cmdprofile_groups())
	profileCmd.AddCommand(profile_users.Cmdprofile_users())
	profileCmd.Short = utils.GenerateCustomDescription(profileCmd.Short, profile_groups.Description, profile_users.Description, )
	profileCmd.Long = profileCmd.Short
}
