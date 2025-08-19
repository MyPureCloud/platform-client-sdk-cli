package workforcemanagement_timeofflimits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeofflimits_available"
)

func init() {
	workforcemanagement_timeofflimitsCmd.AddCommand(workforcemanagement_timeofflimits_available.Cmdworkforcemanagement_timeofflimits_available())
	workforcemanagement_timeofflimitsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_timeofflimitsCmd.Short, workforcemanagement_timeofflimits_available.Description, )
	workforcemanagement_timeofflimitsCmd.Long = workforcemanagement_timeofflimitsCmd.Short
}
