package workforcemanagement_agentschedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agentschedules_managementunits"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agentschedules_mine"
)

func init() {
	workforcemanagement_agentschedulesCmd.AddCommand(workforcemanagement_agentschedules_managementunits.Cmdworkforcemanagement_agentschedules_managementunits())
	workforcemanagement_agentschedulesCmd.AddCommand(workforcemanagement_agentschedules_mine.Cmdworkforcemanagement_agentschedules_mine())
	workforcemanagement_agentschedulesCmd.Short = utils.GenerateCustomDescription(workforcemanagement_agentschedulesCmd.Short, workforcemanagement_agentschedules_managementunits.Description, workforcemanagement_agentschedules_mine.Description, )
	workforcemanagement_agentschedulesCmd.Long = workforcemanagement_agentschedulesCmd.Short
}
