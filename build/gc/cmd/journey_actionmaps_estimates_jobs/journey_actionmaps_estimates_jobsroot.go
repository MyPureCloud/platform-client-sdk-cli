package journey_actionmaps_estimates_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_actionmaps_estimates_jobs_results"
)

func init() {
	journey_actionmaps_estimates_jobsCmd.AddCommand(journey_actionmaps_estimates_jobs_results.Cmdjourney_actionmaps_estimates_jobs_results())
	journey_actionmaps_estimates_jobsCmd.Short = utils.GenerateCustomDescription(journey_actionmaps_estimates_jobsCmd.Short, journey_actionmaps_estimates_jobs_results.Description, )
	journey_actionmaps_estimates_jobsCmd.Long = journey_actionmaps_estimates_jobsCmd.Short
}
