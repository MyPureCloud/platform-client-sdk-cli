package routing_email_domains

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_domains_testconnection"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_domains_validate"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_domains_routes"
)

func init() {
	routing_email_domainsCmd.AddCommand(routing_email_domains_testconnection.Cmdrouting_email_domains_testconnection())
	routing_email_domainsCmd.AddCommand(routing_email_domains_validate.Cmdrouting_email_domains_validate())
	routing_email_domainsCmd.AddCommand(routing_email_domains_routes.Cmdrouting_email_domains_routes())
	routing_email_domainsCmd.Short = utils.GenerateCustomDescription(routing_email_domainsCmd.Short, routing_email_domains_testconnection.Description, routing_email_domains_validate.Description, routing_email_domains_routes.Description, )
	routing_email_domainsCmd.Long = routing_email_domainsCmd.Short
}
