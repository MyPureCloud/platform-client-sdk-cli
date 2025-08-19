package workforcemanagement_agents_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_agents_integrations_hris"
)

func init() {
	workforcemanagement_agents_integrationsCmd.AddCommand(workforcemanagement_agents_integrations_hris.Cmdworkforcemanagement_agents_integrations_hris())
	workforcemanagement_agents_integrationsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_agents_integrationsCmd.Short, workforcemanagement_agents_integrations_hris.Description, )
	workforcemanagement_agents_integrationsCmd.Long = workforcemanagement_agents_integrationsCmd.Short
}
