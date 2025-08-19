package messaging_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/messaging_settings_default"
)

func init() {
	messaging_settingsCmd.AddCommand(messaging_settings_default.Cmdmessaging_settings_default())
	messaging_settingsCmd.Short = utils.GenerateCustomDescription(messaging_settingsCmd.Short, messaging_settings_default.Description, )
	messaging_settingsCmd.Long = messaging_settingsCmd.Short
}
