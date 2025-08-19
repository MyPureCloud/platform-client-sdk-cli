package workforcemanagement_timeoffbalance

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_timeoffbalance_jobs"
)

func init() {
	workforcemanagement_timeoffbalanceCmd.AddCommand(workforcemanagement_timeoffbalance_jobs.Cmdworkforcemanagement_timeoffbalance_jobs())
	workforcemanagement_timeoffbalanceCmd.Short = utils.GenerateCustomDescription(workforcemanagement_timeoffbalanceCmd.Short, workforcemanagement_timeoffbalance_jobs.Description, )
	workforcemanagement_timeoffbalanceCmd.Long = workforcemanagement_timeoffbalanceCmd.Short
}
