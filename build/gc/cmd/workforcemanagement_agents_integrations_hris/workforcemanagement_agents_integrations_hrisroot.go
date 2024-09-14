package workforcemanagement_agents_integrations_hris

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents_integrations_hris_query"
)

func init() {
	workforcemanagement_agents_integrations_hrisCmd.AddCommand(workforcemanagement_agents_integrations_hris_query.Cmdworkforcemanagement_agents_integrations_hris_query())
	workforcemanagement_agents_integrations_hrisCmd.Short = utils.GenerateCustomDescription(workforcemanagement_agents_integrations_hrisCmd.Short, workforcemanagement_agents_integrations_hris_query.Description, )
	workforcemanagement_agents_integrations_hrisCmd.Long = workforcemanagement_agents_integrations_hrisCmd.Short
}