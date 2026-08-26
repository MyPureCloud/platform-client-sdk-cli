package workforcemanagement_schedulebids

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_schedulebids_preference"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_schedulebids_schedulesets"
)

func init() {
	workforcemanagement_schedulebidsCmd.AddCommand(workforcemanagement_schedulebids_preference.Cmdworkforcemanagement_schedulebids_preference())
	workforcemanagement_schedulebidsCmd.AddCommand(workforcemanagement_schedulebids_schedulesets.Cmdworkforcemanagement_schedulebids_schedulesets())
	workforcemanagement_schedulebidsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_schedulebidsCmd.Short, workforcemanagement_schedulebids_preference.Description, workforcemanagement_schedulebids_schedulesets.Description, )
	workforcemanagement_schedulebidsCmd.Long = workforcemanagement_schedulebidsCmd.Short
}
