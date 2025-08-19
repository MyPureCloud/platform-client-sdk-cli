package routing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_queues"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_availablemediatypes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_languages"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_settings"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_skills"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_utilization"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_assessments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_conversations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_message"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_sms"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_predictors"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_skillgroups"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_directroutingbackup"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_users"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_wrapupcodes"
)

func init() {
	routingCmd.AddCommand(routing_queues.Cmdrouting_queues())
	routingCmd.AddCommand(routing_availablemediatypes.Cmdrouting_availablemediatypes())
	routingCmd.AddCommand(routing_languages.Cmdrouting_languages())
	routingCmd.AddCommand(routing_settings.Cmdrouting_settings())
	routingCmd.AddCommand(routing_skills.Cmdrouting_skills())
	routingCmd.AddCommand(routing_utilization.Cmdrouting_utilization())
	routingCmd.AddCommand(routing_assessments.Cmdrouting_assessments())
	routingCmd.AddCommand(routing_conversations.Cmdrouting_conversations())
	routingCmd.AddCommand(routing_email.Cmdrouting_email())
	routingCmd.AddCommand(routing_message.Cmdrouting_message())
	routingCmd.AddCommand(routing_sms.Cmdrouting_sms())
	routingCmd.AddCommand(routing_predictors.Cmdrouting_predictors())
	routingCmd.AddCommand(routing_skillgroups.Cmdrouting_skillgroups())
	routingCmd.AddCommand(routing_directroutingbackup.Cmdrouting_directroutingbackup())
	routingCmd.AddCommand(routing_users.Cmdrouting_users())
	routingCmd.AddCommand(routing_wrapupcodes.Cmdrouting_wrapupcodes())
	routingCmd.Short = utils.GenerateCustomDescription(routingCmd.Short, routing_queues.Description, routing_availablemediatypes.Description, routing_languages.Description, routing_settings.Description, routing_skills.Description, routing_utilization.Description, routing_assessments.Description, routing_conversations.Description, routing_email.Description, routing_message.Description, routing_sms.Description, routing_predictors.Description, routing_skillgroups.Description, routing_directroutingbackup.Description, routing_users.Description, routing_wrapupcodes.Description, )
	routingCmd.Long = routingCmd.Short
}
