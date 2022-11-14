package workforcemanagement_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_adherence_explanations"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_adherence_historical"
)

func init() {
	workforcemanagement_adherenceCmd.AddCommand(workforcemanagement_adherence_explanations.Cmdworkforcemanagement_adherence_explanations())
	workforcemanagement_adherenceCmd.AddCommand(workforcemanagement_adherence_historical.Cmdworkforcemanagement_adherence_historical())
	workforcemanagement_adherenceCmd.Short = utils.GenerateCustomDescription(workforcemanagement_adherenceCmd.Short, workforcemanagement_adherence_explanations.Description, workforcemanagement_adherence_historical.Description, )
	workforcemanagement_adherenceCmd.Long = workforcemanagement_adherenceCmd.Short
}
