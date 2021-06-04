package users_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_me_password"
)

func init() {
	users_meCmd.AddCommand(users_me_password.Cmdusers_me_password())
	users_meCmd.Short = utils.GenerateCustomDescription(users_meCmd.Short, users_me_password.Description, )
	users_meCmd.Long = users_meCmd.Short
}
