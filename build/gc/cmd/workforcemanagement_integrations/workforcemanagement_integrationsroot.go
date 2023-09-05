package workforcemanagement_integrations

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_integrations_hris"
)

func init() {
	workforcemanagement_integrationsCmd.AddCommand(workforcemanagement_integrations_hris.Cmdworkforcemanagement_integrations_hris())
	workforcemanagement_integrationsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_integrationsCmd.Short, workforcemanagement_integrations_hris.Description, )
	workforcemanagement_integrationsCmd.Long = workforcemanagement_integrationsCmd.Short
}
