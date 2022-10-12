package routing_email_outbound

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_outbound_domains"
)

func init() {
	routing_email_outboundCmd.AddCommand(routing_email_outbound_domains.Cmdrouting_email_outbound_domains())
	routing_email_outboundCmd.Short = utils.GenerateCustomDescription(routing_email_outboundCmd.Short, routing_email_outbound_domains.Description, )
	routing_email_outboundCmd.Long = routing_email_outboundCmd.Short
}
