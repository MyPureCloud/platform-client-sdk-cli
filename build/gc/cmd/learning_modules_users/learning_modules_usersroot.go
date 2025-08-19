package learning_modules_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/learning_modules_users_assignments"
)

func init() {
	learning_modules_usersCmd.AddCommand(learning_modules_users_assignments.Cmdlearning_modules_users_assignments())
	learning_modules_usersCmd.Short = utils.GenerateCustomDescription(learning_modules_usersCmd.Short, learning_modules_users_assignments.Description, )
	learning_modules_usersCmd.Long = learning_modules_usersCmd.Short
}
