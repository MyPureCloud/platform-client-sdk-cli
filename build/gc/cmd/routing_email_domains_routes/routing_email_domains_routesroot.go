package routing_email_domains_routes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/routing_email_domains_routes_identityresolution"
)

func init() {
	routing_email_domains_routesCmd.AddCommand(routing_email_domains_routes_identityresolution.Cmdrouting_email_domains_routes_identityresolution())
	routing_email_domains_routesCmd.Short = utils.GenerateCustomDescription(routing_email_domains_routesCmd.Short, routing_email_domains_routes_identityresolution.Description, )
	routing_email_domains_routesCmd.Long = routing_email_domains_routesCmd.Short
}
