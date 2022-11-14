package workforcemanagement_shrinkage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_shrinkage_jobs"
)

func init() {
	workforcemanagement_shrinkageCmd.AddCommand(workforcemanagement_shrinkage_jobs.Cmdworkforcemanagement_shrinkage_jobs())
	workforcemanagement_shrinkageCmd.Short = utils.GenerateCustomDescription(workforcemanagement_shrinkageCmd.Short, workforcemanagement_shrinkage_jobs.Description, )
	workforcemanagement_shrinkageCmd.Long = workforcemanagement_shrinkageCmd.Short
}
