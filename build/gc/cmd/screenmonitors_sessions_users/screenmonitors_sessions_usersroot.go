package screenmonitors_sessions_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/screenmonitors_sessions_users_details"
)

func init() {
	screenmonitors_sessions_usersCmd.AddCommand(screenmonitors_sessions_users_details.Cmdscreenmonitors_sessions_users_details())
	screenmonitors_sessions_usersCmd.Short = utils.GenerateCustomDescription(screenmonitors_sessions_usersCmd.Short, screenmonitors_sessions_users_details.Description, )
	screenmonitors_sessions_usersCmd.Long = screenmonitors_sessions_usersCmd.Short
}
