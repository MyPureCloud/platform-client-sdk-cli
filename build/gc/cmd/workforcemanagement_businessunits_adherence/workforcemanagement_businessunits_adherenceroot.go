package workforcemanagement_businessunits_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_businessunits_adherence_explanations"
)

func init() {
	workforcemanagement_businessunits_adherenceCmd.AddCommand(workforcemanagement_businessunits_adherence_explanations.Cmdworkforcemanagement_businessunits_adherence_explanations())
	workforcemanagement_businessunits_adherenceCmd.Short = utils.GenerateCustomDescription(workforcemanagement_businessunits_adherenceCmd.Short, workforcemanagement_businessunits_adherence_explanations.Description, )
	workforcemanagement_businessunits_adherenceCmd.Long = workforcemanagement_businessunits_adherenceCmd.Short
}
