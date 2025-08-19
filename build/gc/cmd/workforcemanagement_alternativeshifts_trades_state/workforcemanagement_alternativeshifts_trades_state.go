package workforcemanagement_alternativeshifts_trades_state

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_alternativeshifts_trades_state", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/alternativeshifts/trades/state")
	workforcemanagement_alternativeshifts_trades_stateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_alternativeshifts_trades_state"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_alternativeshifts_trades_stateCmd)
}

func Cmdworkforcemanagement_alternativeshifts_trades_state() *cobra.Command {
	return workforcemanagement_alternativeshifts_trades_stateCmd
}
