package journey_actionmaps

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_actionmaps_estimates"
)

func init() {
	journey_actionmapsCmd.AddCommand(journey_actionmaps_estimates.Cmdjourney_actionmaps_estimates())
	journey_actionmapsCmd.Short = utils.GenerateCustomDescription(journey_actionmapsCmd.Short, journey_actionmaps_estimates.Description, )
	journey_actionmapsCmd.Long = journey_actionmapsCmd.Short
}
