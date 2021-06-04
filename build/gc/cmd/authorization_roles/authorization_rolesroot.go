package authorization_roles

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_roles_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_roles_comparedefault"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_roles_subjectgrants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/authorization_roles_default"
)

func init() {
	authorization_rolesCmd.AddCommand(authorization_roles_users.Cmdauthorization_roles_users())
	authorization_rolesCmd.AddCommand(authorization_roles_comparedefault.Cmdauthorization_roles_comparedefault())
	authorization_rolesCmd.AddCommand(authorization_roles_subjectgrants.Cmdauthorization_roles_subjectgrants())
	authorization_rolesCmd.AddCommand(authorization_roles_default.Cmdauthorization_roles_default())
	authorization_rolesCmd.Short = utils.GenerateCustomDescription(authorization_rolesCmd.Short, authorization_roles_users.Description, authorization_roles_comparedefault.Description, authorization_roles_subjectgrants.Description, authorization_roles_default.Description, )
	authorization_rolesCmd.Long = authorization_rolesCmd.Short
}
