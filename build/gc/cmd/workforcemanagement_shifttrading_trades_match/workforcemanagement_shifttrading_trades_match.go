package workforcemanagement_shifttrading_trades_match

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_shifttrading_trades_match", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/shifttrading/trades/{tradeId}/match")
	workforcemanagement_shifttrading_trades_matchCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_shifttrading_trades_match"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_shifttrading_trades_matchCmd)
}

func Cmdworkforcemanagement_shifttrading_trades_match() *cobra.Command {
	return workforcemanagement_shifttrading_trades_matchCmd
}
