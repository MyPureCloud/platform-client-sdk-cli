package routing_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_users_directroutingbackup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_users_utilization"
)

func init() {
	routing_usersCmd.AddCommand(routing_users_directroutingbackup.Cmdrouting_users_directroutingbackup())
	routing_usersCmd.AddCommand(routing_users_utilization.Cmdrouting_users_utilization())
	routing_usersCmd.Short = utils.GenerateCustomDescription(routing_usersCmd.Short, routing_users_directroutingbackup.Description, routing_users_utilization.Description, )
	routing_usersCmd.Long = routing_usersCmd.Short
}
