package workforcemanagement_timeoffrequests

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffrequests_waitlistpositions"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffrequests_integrationstatus"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffrequests_estimate"
)

func init() {
	workforcemanagement_timeoffrequestsCmd.AddCommand(workforcemanagement_timeoffrequests_waitlistpositions.Cmdworkforcemanagement_timeoffrequests_waitlistpositions())
	workforcemanagement_timeoffrequestsCmd.AddCommand(workforcemanagement_timeoffrequests_integrationstatus.Cmdworkforcemanagement_timeoffrequests_integrationstatus())
	workforcemanagement_timeoffrequestsCmd.AddCommand(workforcemanagement_timeoffrequests_estimate.Cmdworkforcemanagement_timeoffrequests_estimate())
	workforcemanagement_timeoffrequestsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_timeoffrequestsCmd.Short, workforcemanagement_timeoffrequests_waitlistpositions.Description, workforcemanagement_timeoffrequests_integrationstatus.Description, workforcemanagement_timeoffrequests_estimate.Description, )
	workforcemanagement_timeoffrequestsCmd.Long = workforcemanagement_timeoffrequestsCmd.Short
}
