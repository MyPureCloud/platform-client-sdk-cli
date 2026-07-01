package screenmonitors_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/screenmonitors_users_sessions"
)

func init() {
	screenmonitors_usersCmd.AddCommand(screenmonitors_users_sessions.Cmdscreenmonitors_users_sessions())
	screenmonitors_usersCmd.Short = utils.GenerateCustomDescription(screenmonitors_usersCmd.Short, screenmonitors_users_sessions.Description, )
	screenmonitors_usersCmd.Long = screenmonitors_usersCmd.Short
}
