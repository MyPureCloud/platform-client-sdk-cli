package workforcemanagement_managementunits_schedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_managementunits_schedules_search"
)

func init() {
	workforcemanagement_managementunits_schedulesCmd.AddCommand(workforcemanagement_managementunits_schedules_search.Cmdworkforcemanagement_managementunits_schedules_search())
	workforcemanagement_managementunits_schedulesCmd.Short = utils.GenerateCustomDescription(workforcemanagement_managementunits_schedulesCmd.Short, workforcemanagement_managementunits_schedules_search.Description, )
	workforcemanagement_managementunits_schedulesCmd.Long = workforcemanagement_managementunits_schedulesCmd.Short
}