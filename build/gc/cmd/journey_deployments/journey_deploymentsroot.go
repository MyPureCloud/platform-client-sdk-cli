package journey_deployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_deployments_actionevent"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_deployments_appevents"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_deployments_customers"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_deployments_webevents"
)

func init() {
	journey_deploymentsCmd.AddCommand(journey_deployments_actionevent.Cmdjourney_deployments_actionevent())
	journey_deploymentsCmd.AddCommand(journey_deployments_appevents.Cmdjourney_deployments_appevents())
	journey_deploymentsCmd.AddCommand(journey_deployments_customers.Cmdjourney_deployments_customers())
	journey_deploymentsCmd.AddCommand(journey_deployments_webevents.Cmdjourney_deployments_webevents())
	journey_deploymentsCmd.Short = utils.GenerateCustomDescription(journey_deploymentsCmd.Short, journey_deployments_actionevent.Description, journey_deployments_appevents.Description, journey_deployments_customers.Description, journey_deployments_webevents.Description, )
	journey_deploymentsCmd.Long = journey_deploymentsCmd.Short
}
