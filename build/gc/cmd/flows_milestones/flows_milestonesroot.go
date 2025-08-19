package flows_milestones

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/flows_milestones_divisionviews"
)

func init() {
	flows_milestonesCmd.AddCommand(flows_milestones_divisionviews.Cmdflows_milestones_divisionviews())
	flows_milestonesCmd.Short = utils.GenerateCustomDescription(flows_milestonesCmd.Short, flows_milestones_divisionviews.Description, )
	flows_milestonesCmd.Long = flows_milestonesCmd.Short
}
