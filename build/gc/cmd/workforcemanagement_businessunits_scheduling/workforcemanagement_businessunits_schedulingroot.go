package workforcemanagement_businessunits_scheduling

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_scheduling_runs"
)

func init() {
	workforcemanagement_businessunits_schedulingCmd.AddCommand(workforcemanagement_businessunits_scheduling_runs.Cmdworkforcemanagement_businessunits_scheduling_runs())
	workforcemanagement_businessunits_schedulingCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_schedulingCmd.Short, workforcemanagement_businessunits_scheduling_runs.Description, )
	workforcemanagement_businessunits_schedulingCmd.Long = workforcemanagement_businessunits_schedulingCmd.Short
}