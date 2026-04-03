package journey_externalevents_configurations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_externalevents_configurations_events"
)

func init() {
	journey_externalevents_configurationsCmd.AddCommand(journey_externalevents_configurations_events.Cmdjourney_externalevents_configurations_events())
	journey_externalevents_configurationsCmd.Short = utils.GenerateCustomDescription(journey_externalevents_configurationsCmd.Short, journey_externalevents_configurations_events.Description, )
	journey_externalevents_configurationsCmd.Long = journey_externalevents_configurationsCmd.Short
}
