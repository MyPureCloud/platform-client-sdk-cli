package workforcemanagement_businessunits_capacityplans_bulk

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplans_bulk", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/capacityplans/bulk")
	workforcemanagement_businessunits_capacityplans_bulkCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplans_bulk"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_capacityplans_bulkCmd)
}

func Cmdworkforcemanagement_businessunits_capacityplans_bulk() *cobra.Command {
	return workforcemanagement_businessunits_capacityplans_bulkCmd
}
