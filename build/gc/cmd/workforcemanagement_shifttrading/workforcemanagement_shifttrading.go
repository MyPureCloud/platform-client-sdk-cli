package workforcemanagement_shifttrading

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_shifttrading", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/shifttrading")
	workforcemanagement_shifttradingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_shifttrading"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_shifttradingCmd)
}

func Cmdworkforcemanagement_shifttrading() *cobra.Command {
	return workforcemanagement_shifttradingCmd
}
