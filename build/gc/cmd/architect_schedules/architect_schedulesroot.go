package architect_schedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/architect_schedules_divisionviews"
)

func init() {
	architect_schedulesCmd.AddCommand(architect_schedules_divisionviews.Cmdarchitect_schedules_divisionviews())
	architect_schedulesCmd.Short = utils.GenerateCustomDescription(architect_schedulesCmd.Short, architect_schedules_divisionviews.Description, )
	architect_schedulesCmd.Long = architect_schedulesCmd.Short
}
