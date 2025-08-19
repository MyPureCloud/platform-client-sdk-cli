package workforcemanagement_agents_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents_me_managementunit"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents_me_possibleworkshifts"
)

func init() {
	workforcemanagement_agents_meCmd.AddCommand(workforcemanagement_agents_me_managementunit.Cmdworkforcemanagement_agents_me_managementunit())
	workforcemanagement_agents_meCmd.AddCommand(workforcemanagement_agents_me_possibleworkshifts.Cmdworkforcemanagement_agents_me_possibleworkshifts())
	workforcemanagement_agents_meCmd.Short = utils.GenerateCustomDescription(workforcemanagement_agents_meCmd.Short, workforcemanagement_agents_me_managementunit.Description, workforcemanagement_agents_me_possibleworkshifts.Description, )
	workforcemanagement_agents_meCmd.Long = workforcemanagement_agents_meCmd.Short
}
