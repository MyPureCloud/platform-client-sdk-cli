package routing_queues

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_estimatedwaittime"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_mediatypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_divisionviews"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_wrapupcodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_members"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_assistant"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_comparisonperiods"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues_identityresolution"
)

func init() {
	routing_queuesCmd.AddCommand(routing_queues_estimatedwaittime.Cmdrouting_queues_estimatedwaittime())
	routing_queuesCmd.AddCommand(routing_queues_mediatypes.Cmdrouting_queues_mediatypes())
	routing_queuesCmd.AddCommand(routing_queues_divisionviews.Cmdrouting_queues_divisionviews())
	routing_queuesCmd.AddCommand(routing_queues_wrapupcodes.Cmdrouting_queues_wrapupcodes())
	routing_queuesCmd.AddCommand(routing_queues_me.Cmdrouting_queues_me())
	routing_queuesCmd.AddCommand(routing_queues_members.Cmdrouting_queues_members())
	routing_queuesCmd.AddCommand(routing_queues_users.Cmdrouting_queues_users())
	routing_queuesCmd.AddCommand(routing_queues_assistant.Cmdrouting_queues_assistant())
	routing_queuesCmd.AddCommand(routing_queues_comparisonperiods.Cmdrouting_queues_comparisonperiods())
	routing_queuesCmd.AddCommand(routing_queues_identityresolution.Cmdrouting_queues_identityresolution())
	routing_queuesCmd.Short = utils.GenerateCustomDescription(routing_queuesCmd.Short, routing_queues_estimatedwaittime.Description, routing_queues_mediatypes.Description, routing_queues_divisionviews.Description, routing_queues_wrapupcodes.Description, routing_queues_me.Description, routing_queues_members.Description, routing_queues_users.Description, routing_queues_assistant.Description, routing_queues_comparisonperiods.Description, routing_queues_identityresolution.Description, )
	routing_queuesCmd.Long = routing_queuesCmd.Short
}
