package workforcemanagement_historicaldata_bulk_remove

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_historicaldata_bulk_remove", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/historicaldata/bulk/remove")
	workforcemanagement_historicaldata_bulk_removeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_historicaldata_bulk_remove"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_historicaldata_bulk_removeCmd)
}

func Cmdworkforcemanagement_historicaldata_bulk_remove() *cobra.Command {
	return workforcemanagement_historicaldata_bulk_removeCmd
}
