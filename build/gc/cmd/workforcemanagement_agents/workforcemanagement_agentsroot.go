package workforcemanagement_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents_me"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents_managementunit"
)

func init() {
	workforcemanagement_agentsCmd.AddCommand(workforcemanagement_agents_me.Cmdworkforcemanagement_agents_me())
	workforcemanagement_agentsCmd.AddCommand(workforcemanagement_agents_managementunit.Cmdworkforcemanagement_agents_managementunit())
	workforcemanagement_agentsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_agentsCmd.Short, workforcemanagement_agents_me.Description, workforcemanagement_agents_managementunit.Description, )
	workforcemanagement_agentsCmd.Long = workforcemanagement_agentsCmd.Short
}
