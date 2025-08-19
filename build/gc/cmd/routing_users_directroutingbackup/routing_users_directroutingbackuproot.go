package routing_users_directroutingbackup

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_users_directroutingbackup_settings"
)

func init() {
	routing_users_directroutingbackupCmd.AddCommand(routing_users_directroutingbackup_settings.Cmdrouting_users_directroutingbackup_settings())
	routing_users_directroutingbackupCmd.Short = utils.GenerateCustomDescription(routing_users_directroutingbackupCmd.Short, routing_users_directroutingbackup_settings.Description, )
	routing_users_directroutingbackupCmd.Long = routing_users_directroutingbackupCmd.Short
}
