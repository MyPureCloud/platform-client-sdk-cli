package workforcemanagement_alternativeshifts_trades_state

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_alternativeshifts_trades_state_jobs"
)

func init() {
	workforcemanagement_alternativeshifts_trades_stateCmd.AddCommand(workforcemanagement_alternativeshifts_trades_state_jobs.Cmdworkforcemanagement_alternativeshifts_trades_state_jobs())
	workforcemanagement_alternativeshifts_trades_stateCmd.Short = utils.GenerateCustomDescription(workforcemanagement_alternativeshifts_trades_stateCmd.Short, workforcemanagement_alternativeshifts_trades_state_jobs.Description, )
	workforcemanagement_alternativeshifts_trades_stateCmd.Long = workforcemanagement_alternativeshifts_trades_stateCmd.Short
}
