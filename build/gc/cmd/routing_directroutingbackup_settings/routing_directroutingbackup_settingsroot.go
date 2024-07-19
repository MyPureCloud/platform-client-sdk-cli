package routing_directroutingbackup_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_directroutingbackup_settings_me"
)

func init() {
	routing_directroutingbackup_settingsCmd.AddCommand(routing_directroutingbackup_settings_me.Cmdrouting_directroutingbackup_settings_me())
	routing_directroutingbackup_settingsCmd.Short = utils.GenerateCustomDescription(routing_directroutingbackup_settingsCmd.Short, routing_directroutingbackup_settings_me.Description, )
	routing_directroutingbackup_settingsCmd.Long = routing_directroutingbackup_settingsCmd.Short
}
