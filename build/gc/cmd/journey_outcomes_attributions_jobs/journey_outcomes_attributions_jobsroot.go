package journey_outcomes_attributions_jobs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_outcomes_attributions_jobs_results"
)

func init() {
	journey_outcomes_attributions_jobsCmd.AddCommand(journey_outcomes_attributions_jobs_results.Cmdjourney_outcomes_attributions_jobs_results())
	journey_outcomes_attributions_jobsCmd.Short = utils.GenerateCustomDescription(journey_outcomes_attributions_jobsCmd.Short, journey_outcomes_attributions_jobs_results.Description, )
	journey_outcomes_attributions_jobsCmd.Long = journey_outcomes_attributions_jobsCmd.Short
}
