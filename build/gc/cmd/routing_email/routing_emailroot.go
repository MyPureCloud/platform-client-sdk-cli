package routing_email

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_domains"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_outbound"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_setup"
)

func init() {
	routing_emailCmd.AddCommand(routing_email_domains.Cmdrouting_email_domains())
	routing_emailCmd.AddCommand(routing_email_outbound.Cmdrouting_email_outbound())
	routing_emailCmd.AddCommand(routing_email_setup.Cmdrouting_email_setup())
	routing_emailCmd.Short = utils.GenerateCustomDescription(routing_emailCmd.Short, routing_email_domains.Description, routing_email_outbound.Description, routing_email_setup.Description, )
	routing_emailCmd.Long = routing_emailCmd.Short
}
