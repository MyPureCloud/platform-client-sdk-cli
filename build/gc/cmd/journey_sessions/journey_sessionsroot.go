package journey_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_sessions_outcomescores"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_sessions_events"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_sessions_segments"
)

func init() {
	journey_sessionsCmd.AddCommand(journey_sessions_outcomescores.Cmdjourney_sessions_outcomescores())
	journey_sessionsCmd.AddCommand(journey_sessions_events.Cmdjourney_sessions_events())
	journey_sessionsCmd.AddCommand(journey_sessions_segments.Cmdjourney_sessions_segments())
	journey_sessionsCmd.Short = utils.GenerateCustomDescription(journey_sessionsCmd.Short, journey_sessions_outcomescores.Description, journey_sessions_events.Description, journey_sessions_segments.Description, )
	journey_sessionsCmd.Long = journey_sessionsCmd.Short
}
