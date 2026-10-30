package workforcemanagement_selfschedule_offers

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_selfschedule_offers_jobs"
)

func init() {
	workforcemanagement_selfschedule_offersCmd.AddCommand(workforcemanagement_selfschedule_offers_jobs.Cmdworkforcemanagement_selfschedule_offers_jobs())
	workforcemanagement_selfschedule_offersCmd.Short = utils.GenerateCustomDescription(workforcemanagement_selfschedule_offersCmd.Short, workforcemanagement_selfschedule_offers_jobs.Description, )
	workforcemanagement_selfschedule_offersCmd.Long = workforcemanagement_selfschedule_offersCmd.Short
}
