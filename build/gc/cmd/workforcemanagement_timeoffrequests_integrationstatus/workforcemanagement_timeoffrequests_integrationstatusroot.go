package workforcemanagement_timeoffrequests_integrationstatus

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffrequests_integrationstatus_query"
)

func init() {
	workforcemanagement_timeoffrequests_integrationstatusCmd.AddCommand(workforcemanagement_timeoffrequests_integrationstatus_query.Cmdworkforcemanagement_timeoffrequests_integrationstatus_query())
	workforcemanagement_timeoffrequests_integrationstatusCmd.Short = utils.GenerateCustomDescription(workforcemanagement_timeoffrequests_integrationstatusCmd.Short, workforcemanagement_timeoffrequests_integrationstatus_query.Description, )
	workforcemanagement_timeoffrequests_integrationstatusCmd.Long = workforcemanagement_timeoffrequests_integrationstatusCmd.Short
}
