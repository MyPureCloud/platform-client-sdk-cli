package routing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_wrapupcodes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_message"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_languages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_sms"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_skills"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_utilization"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_conversations"
)

func init() {
	routingCmd.AddCommand(routing_email.Cmdrouting_email())
	routingCmd.AddCommand(routing_wrapupcodes.Cmdrouting_wrapupcodes())
	routingCmd.AddCommand(routing_queues.Cmdrouting_queues())
	routingCmd.AddCommand(routing_message.Cmdrouting_message())
	routingCmd.AddCommand(routing_languages.Cmdrouting_languages())
	routingCmd.AddCommand(routing_users.Cmdrouting_users())
	routingCmd.AddCommand(routing_sms.Cmdrouting_sms())
	routingCmd.AddCommand(routing_settings.Cmdrouting_settings())
	routingCmd.AddCommand(routing_skills.Cmdrouting_skills())
	routingCmd.AddCommand(routing_utilization.Cmdrouting_utilization())
	routingCmd.AddCommand(routing_conversations.Cmdrouting_conversations())
	routingCmd.Short = utils.GenerateCustomDescription(routingCmd.Short, routing_email.Description, routing_wrapupcodes.Description, routing_queues.Description, routing_message.Description, routing_languages.Description, routing_users.Description, routing_sms.Description, routing_settings.Description, routing_skills.Description, routing_utilization.Description, routing_conversations.Description, )
	routingCmd.Long = routingCmd.Short
}
