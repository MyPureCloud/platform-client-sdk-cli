package routing_utilization_tags

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_utilization_tags_agents"
)

func init() {
	routing_utilization_tagsCmd.AddCommand(routing_utilization_tags_agents.Cmdrouting_utilization_tags_agents())
	routing_utilization_tagsCmd.Short = utils.GenerateCustomDescription(routing_utilization_tagsCmd.Short, routing_utilization_tags_agents.Description, )
	routing_utilization_tagsCmd.Long = routing_utilization_tagsCmd.Short
}
