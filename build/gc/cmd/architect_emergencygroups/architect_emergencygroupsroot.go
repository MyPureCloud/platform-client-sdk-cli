package architect_emergencygroups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_emergencygroups_divisionviews"
)

func init() {
	architect_emergencygroupsCmd.AddCommand(architect_emergencygroups_divisionviews.Cmdarchitect_emergencygroups_divisionviews())
	architect_emergencygroupsCmd.Short = utils.GenerateCustomDescription(architect_emergencygroupsCmd.Short, architect_emergencygroups_divisionviews.Description, )
	architect_emergencygroupsCmd.Long = architect_emergencygroupsCmd.Short
}
