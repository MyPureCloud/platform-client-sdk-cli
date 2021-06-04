package routing_queues_divisionviews

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_divisionviews_all"
)

func init() {
	routing_queues_divisionviewsCmd.AddCommand(routing_queues_divisionviews_all.Cmdrouting_queues_divisionviews_all())
	routing_queues_divisionviewsCmd.Short = utils.GenerateCustomDescription(routing_queues_divisionviewsCmd.Short, routing_queues_divisionviews_all.Description, )
	routing_queues_divisionviewsCmd.Long = routing_queues_divisionviewsCmd.Short
}
