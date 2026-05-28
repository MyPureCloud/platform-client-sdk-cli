package routing_skillexpressions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_skillexpressions_validate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_skillexpressions_queue"
)

func init() {
	routing_skillexpressionsCmd.AddCommand(routing_skillexpressions_validate.Cmdrouting_skillexpressions_validate())
	routing_skillexpressionsCmd.AddCommand(routing_skillexpressions_queue.Cmdrouting_skillexpressions_queue())
	routing_skillexpressionsCmd.Short = utils.GenerateCustomDescription(routing_skillexpressionsCmd.Short, routing_skillexpressions_validate.Description, routing_skillexpressions_queue.Description, )
	routing_skillexpressionsCmd.Long = routing_skillexpressionsCmd.Short
}
