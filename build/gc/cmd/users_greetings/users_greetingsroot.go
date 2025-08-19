package users_greetings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_greetings_defaults"
)

func init() {
	users_greetingsCmd.AddCommand(users_greetings_defaults.Cmdusers_greetings_defaults())
	users_greetingsCmd.Short = utils.GenerateCustomDescription(users_greetingsCmd.Short, users_greetings_defaults.Description, )
	users_greetingsCmd.Long = users_greetingsCmd.Short
}
