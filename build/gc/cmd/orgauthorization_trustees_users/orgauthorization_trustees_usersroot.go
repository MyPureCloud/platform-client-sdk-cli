package orgauthorization_trustees_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustees_users_roles"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/orgauthorization_trustees_users_roledivisions"
)

func init() {
	orgauthorization_trustees_usersCmd.AddCommand(orgauthorization_trustees_users_roles.Cmdorgauthorization_trustees_users_roles())
	orgauthorization_trustees_usersCmd.AddCommand(orgauthorization_trustees_users_roledivisions.Cmdorgauthorization_trustees_users_roledivisions())
	orgauthorization_trustees_usersCmd.Short = utils.GenerateCustomDescription(orgauthorization_trustees_usersCmd.Short, orgauthorization_trustees_users_roles.Description, orgauthorization_trustees_users_roledivisions.Description, )
	orgauthorization_trustees_usersCmd.Long = orgauthorization_trustees_usersCmd.Short
}
