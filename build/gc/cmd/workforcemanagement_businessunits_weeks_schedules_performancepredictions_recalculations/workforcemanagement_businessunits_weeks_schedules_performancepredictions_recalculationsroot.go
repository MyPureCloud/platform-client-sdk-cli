package workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations_uploadurl"
)

func init() {
	workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculationsCmd.AddCommand(workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations_uploadurl.Cmdworkforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations_uploadurl())
	workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculationsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculationsCmd.Short, workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculations_uploadurl.Description, )
	workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculationsCmd.Long = workforcemanagement_businessunits_weeks_schedules_performancepredictions_recalculationsCmd.Short
}
