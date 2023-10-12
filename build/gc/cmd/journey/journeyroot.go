package journey

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_actionmaps"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_actiontargets"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_actiontemplates"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_outcomes"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_sessions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_segments"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_deployments"
)

func init() {
	journeyCmd.AddCommand(journey_actionmaps.Cmdjourney_actionmaps())
	journeyCmd.AddCommand(journey_actiontargets.Cmdjourney_actiontargets())
	journeyCmd.AddCommand(journey_actiontemplates.Cmdjourney_actiontemplates())
	journeyCmd.AddCommand(journey_outcomes.Cmdjourney_outcomes())
	journeyCmd.AddCommand(journey_sessions.Cmdjourney_sessions())
	journeyCmd.AddCommand(journey_segments.Cmdjourney_segments())
	journeyCmd.AddCommand(journey_deployments.Cmdjourney_deployments())
	journeyCmd.Short = utils.GenerateCustomDescription(journeyCmd.Short, journey_actionmaps.Description, journey_actiontargets.Description, journey_actiontemplates.Description, journey_outcomes.Description, journey_sessions.Description, journey_segments.Description, journey_deployments.Description, )
	journeyCmd.Long = journeyCmd.Short
}
