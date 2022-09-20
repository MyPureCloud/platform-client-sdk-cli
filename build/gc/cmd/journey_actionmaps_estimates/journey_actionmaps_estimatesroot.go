package journey_actionmaps_estimates

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_actionmaps_estimates_jobs"
)

func init() {
	journey_actionmaps_estimatesCmd.AddCommand(journey_actionmaps_estimates_jobs.Cmdjourney_actionmaps_estimates_jobs())
	journey_actionmaps_estimatesCmd.Short = utils.GenerateCustomDescription(journey_actionmaps_estimatesCmd.Short, journey_actionmaps_estimates_jobs.Description, )
	journey_actionmaps_estimatesCmd.Long = journey_actionmaps_estimatesCmd.Short
}
