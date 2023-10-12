package journey_deployments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_deployments_appevents"
)

func init() {
	journey_deploymentsCmd.AddCommand(journey_deployments_appevents.Cmdjourney_deployments_appevents())
	journey_deploymentsCmd.Short = utils.GenerateCustomDescription(journey_deploymentsCmd.Short, journey_deployments_appevents.Description, )
	journey_deploymentsCmd.Long = journey_deploymentsCmd.Short
}
