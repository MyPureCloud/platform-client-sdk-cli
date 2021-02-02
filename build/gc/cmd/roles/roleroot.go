package roles

import "github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/roleuser"

func init() {
	rolesCmd.AddCommand(roleuser.Cmdroleuser())
}