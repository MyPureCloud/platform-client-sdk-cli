package presence_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/presence_users_primarysource"
)

func init() {
	presence_usersCmd.AddCommand(presence_users_primarysource.Cmdpresence_users_primarysource())
	presence_usersCmd.Short = utils.GenerateCustomDescription(presence_usersCmd.Short, presence_users_primarysource.Description, )
	presence_usersCmd.Long = presence_usersCmd.Short
}
