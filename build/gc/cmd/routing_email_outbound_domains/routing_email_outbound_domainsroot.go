package routing_email_outbound_domains

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_outbound_domains_search"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_outbound_domains_activation"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_outbound_domains_simulated"
)

func init() {
	routing_email_outbound_domainsCmd.AddCommand(routing_email_outbound_domains_search.Cmdrouting_email_outbound_domains_search())
	routing_email_outbound_domainsCmd.AddCommand(routing_email_outbound_domains_activation.Cmdrouting_email_outbound_domains_activation())
	routing_email_outbound_domainsCmd.AddCommand(routing_email_outbound_domains_simulated.Cmdrouting_email_outbound_domains_simulated())
	routing_email_outbound_domainsCmd.Short = utils.GenerateCustomDescription(routing_email_outbound_domainsCmd.Short, routing_email_outbound_domains_search.Description, routing_email_outbound_domains_activation.Description, routing_email_outbound_domains_simulated.Description, )
	routing_email_outbound_domainsCmd.Long = routing_email_outbound_domainsCmd.Short
}
