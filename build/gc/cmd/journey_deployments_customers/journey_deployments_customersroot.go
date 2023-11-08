package journey_deployments_customers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_deployments_customers_ping"
)

func init() {
	journey_deployments_customersCmd.AddCommand(journey_deployments_customers_ping.Cmdjourney_deployments_customers_ping())
	journey_deployments_customersCmd.Short = utils.GenerateCustomDescription(journey_deployments_customersCmd.Short, journey_deployments_customers_ping.Description, )
	journey_deployments_customersCmd.Long = journey_deployments_customersCmd.Short
}
