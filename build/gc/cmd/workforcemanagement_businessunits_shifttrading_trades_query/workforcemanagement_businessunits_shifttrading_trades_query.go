package workforcemanagement_businessunits_shifttrading_trades_query

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_trades_query", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/shifttrading/trades/query")
	workforcemanagement_businessunits_shifttrading_trades_queryCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_trades_query"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_shifttrading_trades_queryCmd)
}

func Cmdworkforcemanagement_businessunits_shifttrading_trades_query() *cobra.Command {
	return workforcemanagement_businessunits_shifttrading_trades_queryCmd
}
