package emails_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/emails_settings_threading"
)

func init() {
	emails_settingsCmd.AddCommand(emails_settings_threading.Cmdemails_settings_threading())
	emails_settingsCmd.Short = utils.GenerateCustomDescription(emails_settingsCmd.Short, emails_settings_threading.Description, )
	emails_settingsCmd.Long = emails_settingsCmd.Short
}
