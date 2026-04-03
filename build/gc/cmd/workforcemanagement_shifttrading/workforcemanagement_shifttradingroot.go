package workforcemanagement_shifttrading

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/cmd/workforcemanagement_shifttrading_trades"
)

func init() {
	workforcemanagement_shifttradingCmd.AddCommand(workforcemanagement_shifttrading_trades.Cmdworkforcemanagement_shifttrading_trades())
	workforcemanagement_shifttradingCmd.Short = utils.GenerateCustomDescription(workforcemanagement_shifttradingCmd.Short, workforcemanagement_shifttrading_trades.Description, )
	workforcemanagement_shifttradingCmd.Long = workforcemanagement_shifttradingCmd.Short
}
