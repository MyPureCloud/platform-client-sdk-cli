package authorization_roles_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_roles_users_add"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_roles_users_remove"
)

func init() {
	authorization_roles_usersCmd.AddCommand(authorization_roles_users_add.Cmdauthorization_roles_users_add())
	authorization_roles_usersCmd.AddCommand(authorization_roles_users_remove.Cmdauthorization_roles_users_remove())
	authorization_roles_usersCmd.Short = utils.GenerateCustomDescription(authorization_roles_usersCmd.Short, authorization_roles_users_add.Description, authorization_roles_users_remove.Description, )
	authorization_roles_usersCmd.Long = authorization_roles_usersCmd.Short
}
