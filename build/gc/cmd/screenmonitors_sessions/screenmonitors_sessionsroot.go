package screenmonitors_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/screenmonitors_sessions_details"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/screenmonitors_sessions_users"
)

func init() {
	screenmonitors_sessionsCmd.AddCommand(screenmonitors_sessions_details.Cmdscreenmonitors_sessions_details())
	screenmonitors_sessionsCmd.AddCommand(screenmonitors_sessions_users.Cmdscreenmonitors_sessions_users())
	screenmonitors_sessionsCmd.Short = utils.GenerateCustomDescription(screenmonitors_sessionsCmd.Short, screenmonitors_sessions_details.Description, screenmonitors_sessions_users.Description, )
	screenmonitors_sessionsCmd.Long = screenmonitors_sessionsCmd.Short
}
