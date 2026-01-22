package workforcemanagement_agentschedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agentschedules_mine"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agentschedules_managementunits"
)

func init() {
	workforcemanagement_agentschedulesCmd.AddCommand(workforcemanagement_agentschedules_mine.Cmdworkforcemanagement_agentschedules_mine())
	workforcemanagement_agentschedulesCmd.AddCommand(workforcemanagement_agentschedules_managementunits.Cmdworkforcemanagement_agentschedules_managementunits())
	workforcemanagement_agentschedulesCmd.Short = utils.GenerateCustomDescription(workforcemanagement_agentschedulesCmd.Short, workforcemanagement_agentschedules_mine.Description, workforcemanagement_agentschedules_managementunits.Description, )
	workforcemanagement_agentschedulesCmd.Long = workforcemanagement_agentschedulesCmd.Short
}
