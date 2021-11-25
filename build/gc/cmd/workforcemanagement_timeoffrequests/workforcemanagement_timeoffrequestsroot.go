package workforcemanagement_timeoffrequests

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffrequests_waitlistpositions"
)

func init() {
	workforcemanagement_timeoffrequestsCmd.AddCommand(workforcemanagement_timeoffrequests_waitlistpositions.Cmdworkforcemanagement_timeoffrequests_waitlistpositions())
	workforcemanagement_timeoffrequestsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_timeoffrequestsCmd.Short, workforcemanagement_timeoffrequests_waitlistpositions.Description, )
	workforcemanagement_timeoffrequestsCmd.Long = workforcemanagement_timeoffrequestsCmd.Short
}
