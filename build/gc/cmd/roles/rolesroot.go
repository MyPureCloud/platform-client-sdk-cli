
package roles

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/user_roles"
)

func init() {
	rolesCmd.AddCommand(user_roles.Cmduser_roles())
}
