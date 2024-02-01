package journey_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_sessions_outcomescores"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_sessions_events"
)

func init() {
	journey_sessionsCmd.AddCommand(journey_sessions_outcomescores.Cmdjourney_sessions_outcomescores())
	journey_sessionsCmd.AddCommand(journey_sessions_events.Cmdjourney_sessions_events())
	journey_sessionsCmd.Short = utils.GenerateCustomDescription(journey_sessionsCmd.Short, journey_sessions_outcomescores.Description, journey_sessions_events.Description, )
	journey_sessionsCmd.Long = journey_sessionsCmd.Short
}
