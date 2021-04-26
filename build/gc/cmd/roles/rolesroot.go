
package roles

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/subjectgrants"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_roles"
)

func init() {
	rolesCmd.AddCommand(subjectgrants.Cmdsubjectgrants())
	rolesCmd.AddCommand(users_roles.Cmdusers_roles())
}
