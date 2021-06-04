package users_routingskills

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/users_routingskills_bulk"
)

func init() {
	users_routingskillsCmd.AddCommand(users_routingskills_bulk.Cmdusers_routingskills_bulk())
	users_routingskillsCmd.Short = utils.GenerateCustomDescription(users_routingskillsCmd.Short, users_routingskills_bulk.Description, )
	users_routingskillsCmd.Long = users_routingskillsCmd.Short
}
