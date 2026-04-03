package workforcemanagement_shifttrading_trades_mine_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_shifttrading_trades_mine_query", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/shifttrading/trades/mine/query")
	workforcemanagement_shifttrading_trades_mine_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_shifttrading_trades_mine_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_shifttrading_trades_mine_queryCmd)
}

func Cmdworkforcemanagement_shifttrading_trades_mine_query() *cobra.Command {
	return workforcemanagement_shifttrading_trades_mine_queryCmd
}
