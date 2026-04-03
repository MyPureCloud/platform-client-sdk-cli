package workforcemanagement_shifttrading_trades

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_shifttrading_trades", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/shifttrading/trades")
	workforcemanagement_shifttrading_tradesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_shifttrading_trades"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_shifttrading_tradesCmd)
}

func Cmdworkforcemanagement_shifttrading_trades() *cobra.Command {
	return workforcemanagement_shifttrading_tradesCmd
}
