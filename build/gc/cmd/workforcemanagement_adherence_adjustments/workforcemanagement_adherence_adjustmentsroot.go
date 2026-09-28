package workforcemanagement_adherence_adjustments

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_adherence_adjustments_query"
)

func init() {
	workforcemanagement_adherence_adjustmentsCmd.AddCommand(workforcemanagement_adherence_adjustments_query.Cmdworkforcemanagement_adherence_adjustments_query())
	workforcemanagement_adherence_adjustmentsCmd.Short = utils.GenerateCustomDescription(workforcemanagement_adherence_adjustmentsCmd.Short, workforcemanagement_adherence_adjustments_query.Description, )
	workforcemanagement_adherence_adjustmentsCmd.Long = workforcemanagement_adherence_adjustmentsCmd.Short
}
