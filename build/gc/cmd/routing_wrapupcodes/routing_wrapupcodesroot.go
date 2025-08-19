package routing_wrapupcodes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_wrapupcodes_divisionviews"
)

func init() {
	routing_wrapupcodesCmd.AddCommand(routing_wrapupcodes_divisionviews.Cmdrouting_wrapupcodes_divisionviews())
	routing_wrapupcodesCmd.Short = utils.GenerateCustomDescription(routing_wrapupcodesCmd.Short, routing_wrapupcodes_divisionviews.Description, )
	routing_wrapupcodesCmd.Long = routing_wrapupcodesCmd.Short
}
