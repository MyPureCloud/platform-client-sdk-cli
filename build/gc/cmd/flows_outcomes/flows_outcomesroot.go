package flows_outcomes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_outcomes_divisionviews"
)

func init() {
	flows_outcomesCmd.AddCommand(flows_outcomes_divisionviews.Cmdflows_outcomes_divisionviews())
	flows_outcomesCmd.Short = utils.GenerateCustomDescription(flows_outcomesCmd.Short, flows_outcomes_divisionviews.Description, )
	flows_outcomesCmd.Long = flows_outcomesCmd.Short
}
