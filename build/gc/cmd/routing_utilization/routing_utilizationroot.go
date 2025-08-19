package routing_utilization

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_utilization_labels"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_utilization_tags"
)

func init() {
	routing_utilizationCmd.AddCommand(routing_utilization_labels.Cmdrouting_utilization_labels())
	routing_utilizationCmd.AddCommand(routing_utilization_tags.Cmdrouting_utilization_tags())
	routing_utilizationCmd.Short = utils.GenerateCustomDescription(routing_utilizationCmd.Short, routing_utilization_labels.Description, routing_utilization_tags.Description, )
	routing_utilizationCmd.Long = routing_utilizationCmd.Short
}
