package routing_utilization_labels

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_utilization_labels_agents"
)

func init() {
	routing_utilization_labelsCmd.AddCommand(routing_utilization_labels_agents.Cmdrouting_utilization_labels_agents())
	routing_utilization_labelsCmd.Short = utils.GenerateCustomDescription(routing_utilization_labelsCmd.Short, routing_utilization_labels_agents.Description, )
	routing_utilization_labelsCmd.Long = routing_utilization_labelsCmd.Short
}
