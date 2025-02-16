package workforcemanagement_historicaldata_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_historicaldata_bulk", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/historicaldata/bulk")
	workforcemanagement_historicaldata_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_historicaldata_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_historicaldata_bulkCmd)
}

func Cmdworkforcemanagement_historicaldata_bulk() *cobra.Command {
	return workforcemanagement_historicaldata_bulkCmd
}
