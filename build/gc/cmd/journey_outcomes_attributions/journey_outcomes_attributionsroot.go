package journey_outcomes_attributions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_outcomes_attributions_jobs"
)

func init() {
	journey_outcomes_attributionsCmd.AddCommand(journey_outcomes_attributions_jobs.Cmdjourney_outcomes_attributions_jobs())
	journey_outcomes_attributionsCmd.Short = utils.GenerateCustomDescription(journey_outcomes_attributionsCmd.Short, journey_outcomes_attributions_jobs.Description, )
	journey_outcomes_attributionsCmd.Long = journey_outcomes_attributionsCmd.Short
}
