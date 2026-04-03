package workforcemanagement_businessunits_shifttrading_trades_evaluate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_trades_evaluate", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/shifttrading/trades/evaluate")
	workforcemanagement_businessunits_shifttrading_trades_evaluateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_trades_evaluate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_shifttrading_trades_evaluateCmd)
}

func Cmdworkforcemanagement_businessunits_shifttrading_trades_evaluate() *cobra.Command {
	return workforcemanagement_businessunits_shifttrading_trades_evaluateCmd
}
