package integrations_workforcemanagement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/integrations_workforcemanagement_vendorconnection"
)

func init() {
	integrations_workforcemanagementCmd.AddCommand(integrations_workforcemanagement_vendorconnection.Cmdintegrations_workforcemanagement_vendorconnection())
	integrations_workforcemanagementCmd.Short = utils.GenerateCustomDescription(integrations_workforcemanagementCmd.Short, integrations_workforcemanagement_vendorconnection.Description, )
	integrations_workforcemanagementCmd.Long = integrations_workforcemanagementCmd.Short
}
