package workforcemanagement_adherence_historical

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_adherence_historical_jobs"
)

func init() {
	workforcemanagement_adherence_historicalCmd.AddCommand(workforcemanagement_adherence_historical_jobs.Cmdworkforcemanagement_adherence_historical_jobs())
	workforcemanagement_adherence_historicalCmd.Short = utils.GenerateCustomDescription(workforcemanagement_adherence_historicalCmd.Short, workforcemanagement_adherence_historical_jobs.Description, )
	workforcemanagement_adherence_historicalCmd.Long = workforcemanagement_adherence_historicalCmd.Short
}
