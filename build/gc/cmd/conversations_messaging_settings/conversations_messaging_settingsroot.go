package conversations_messaging_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/conversations_messaging_settings_default"
)

func init() {
	conversations_messaging_settingsCmd.AddCommand(conversations_messaging_settings_default.Cmdconversations_messaging_settings_default())
	conversations_messaging_settingsCmd.Short = utils.GenerateCustomDescription(conversations_messaging_settingsCmd.Short, conversations_messaging_settings_default.Description, )
	conversations_messaging_settingsCmd.Long = conversations_messaging_settingsCmd.Short
}
