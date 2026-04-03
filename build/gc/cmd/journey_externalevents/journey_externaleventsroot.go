package journey_externalevents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_externalevents_schemas"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_externalevents_configurations"
)

func init() {
	journey_externaleventsCmd.AddCommand(journey_externalevents_schemas.Cmdjourney_externalevents_schemas())
	journey_externaleventsCmd.AddCommand(journey_externalevents_configurations.Cmdjourney_externalevents_configurations())
	journey_externaleventsCmd.Short = utils.GenerateCustomDescription(journey_externaleventsCmd.Short, journey_externalevents_schemas.Description, journey_externalevents_configurations.Description, )
	journey_externaleventsCmd.Long = journey_externaleventsCmd.Short
}
