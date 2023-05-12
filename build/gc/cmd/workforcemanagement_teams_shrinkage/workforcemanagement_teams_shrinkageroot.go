package workforcemanagement_teams_shrinkage

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_teams_shrinkage_jobs"
)

func init() {
	workforcemanagement_teams_shrinkageCmd.AddCommand(workforcemanagement_teams_shrinkage_jobs.Cmdworkforcemanagement_teams_shrinkage_jobs())
	workforcemanagement_teams_shrinkageCmd.Short = utils.GenerateCustomDescription(workforcemanagement_teams_shrinkageCmd.Short, workforcemanagement_teams_shrinkage_jobs.Description, )
	workforcemanagement_teams_shrinkageCmd.Long = workforcemanagement_teams_shrinkageCmd.Short
}
