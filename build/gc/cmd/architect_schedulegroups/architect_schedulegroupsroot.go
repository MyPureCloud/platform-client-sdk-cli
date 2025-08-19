package architect_schedulegroups

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_schedulegroups_divisionviews"
)

func init() {
	architect_schedulegroupsCmd.AddCommand(architect_schedulegroups_divisionviews.Cmdarchitect_schedulegroups_divisionviews())
	architect_schedulegroupsCmd.Short = utils.GenerateCustomDescription(architect_schedulegroupsCmd.Short, architect_schedulegroups_divisionviews.Description, )
	architect_schedulegroupsCmd.Long = architect_schedulegroupsCmd.Short
}
