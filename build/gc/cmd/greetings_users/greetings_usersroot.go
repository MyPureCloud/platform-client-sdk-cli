package greetings_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/greetings_users_downloads"
)

func init() {
	greetings_usersCmd.AddCommand(greetings_users_downloads.Cmdgreetings_users_downloads())
	greetings_usersCmd.Short = utils.GenerateCustomDescription(greetings_usersCmd.Short, greetings_users_downloads.Description, )
	greetings_usersCmd.Long = greetings_usersCmd.Short
}
