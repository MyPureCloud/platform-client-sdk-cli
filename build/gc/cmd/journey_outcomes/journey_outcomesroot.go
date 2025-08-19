package journey_outcomes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_outcomes_predictors"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/journey_outcomes_attributions"
)

func init() {
	journey_outcomesCmd.AddCommand(journey_outcomes_predictors.Cmdjourney_outcomes_predictors())
	journey_outcomesCmd.AddCommand(journey_outcomes_attributions.Cmdjourney_outcomes_attributions())
	journey_outcomesCmd.Short = utils.GenerateCustomDescription(journey_outcomesCmd.Short, journey_outcomes_predictors.Description, journey_outcomes_attributions.Description, )
	journey_outcomesCmd.Long = journey_outcomesCmd.Short
}
