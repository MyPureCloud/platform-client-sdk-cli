package settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/settings_executiondata"
)

func init() {
	settingsCmd.AddCommand(settings_executiondata.Cmdsettings_executiondata())
	settingsCmd.Short = utils.GenerateCustomDescription(settingsCmd.Short, settings_executiondata.Description, )
	settingsCmd.Long = settingsCmd.Short
}
