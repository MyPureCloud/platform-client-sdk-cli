package screenmonitors

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/screenmonitors_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/screenmonitors_sessions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/screenmonitors_settings"
)

func init() {
	screenmonitorsCmd.AddCommand(screenmonitors_users.Cmdscreenmonitors_users())
	screenmonitorsCmd.AddCommand(screenmonitors_sessions.Cmdscreenmonitors_sessions())
	screenmonitorsCmd.AddCommand(screenmonitors_settings.Cmdscreenmonitors_settings())
	screenmonitorsCmd.Short = utils.GenerateCustomDescription(screenmonitorsCmd.Short, screenmonitors_users.Description, screenmonitors_sessions.Description, screenmonitors_settings.Description, )
	screenmonitorsCmd.Long = screenmonitorsCmd.Short
}
