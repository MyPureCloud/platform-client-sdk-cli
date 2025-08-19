package workforcemanagement_integrations_hris

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_integrations_hris_timeofftypes"
)

func init() {
	workforcemanagement_integrations_hrisCmd.AddCommand(workforcemanagement_integrations_hris_timeofftypes.Cmdworkforcemanagement_integrations_hris_timeofftypes())
	workforcemanagement_integrations_hrisCmd.Short = utils.GenerateCustomDescription(workforcemanagement_integrations_hrisCmd.Short, workforcemanagement_integrations_hris_timeofftypes.Description, )
	workforcemanagement_integrations_hrisCmd.Long = workforcemanagement_integrations_hrisCmd.Short
}
