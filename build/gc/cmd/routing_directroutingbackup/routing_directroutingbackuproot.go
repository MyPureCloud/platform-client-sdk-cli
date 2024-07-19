package routing_directroutingbackup

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_directroutingbackup_settings"
)

func init() {
	routing_directroutingbackupCmd.AddCommand(routing_directroutingbackup_settings.Cmdrouting_directroutingbackup_settings())
	routing_directroutingbackupCmd.Short = utils.GenerateCustomDescription(routing_directroutingbackupCmd.Short, routing_directroutingbackup_settings.Description, )
	routing_directroutingbackupCmd.Long = routing_directroutingbackupCmd.Short
}
