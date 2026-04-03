package workforcemanagement_shifttrading_trades_mine

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_shifttrading_trades_mine", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/shifttrading/trades/mine")
	workforcemanagement_shifttrading_trades_mineCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_shifttrading_trades_mine"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_shifttrading_trades_mineCmd)
}

func Cmdworkforcemanagement_shifttrading_trades_mine() *cobra.Command {
	return workforcemanagement_shifttrading_trades_mineCmd
}
