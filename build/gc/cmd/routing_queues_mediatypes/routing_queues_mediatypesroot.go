package routing_queues_mediatypes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_mediatypes_estimatedwaittime"
)

func init() {
	routing_queues_mediatypesCmd.AddCommand(routing_queues_mediatypes_estimatedwaittime.Cmdrouting_queues_mediatypes_estimatedwaittime())
	routing_queues_mediatypesCmd.Short = utils.GenerateCustomDescription(routing_queues_mediatypesCmd.Short, routing_queues_mediatypes_estimatedwaittime.Description, )
	routing_queues_mediatypesCmd.Long = routing_queues_mediatypesCmd.Short
}
