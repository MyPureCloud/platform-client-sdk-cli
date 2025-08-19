package workforcemanagement_businessunits_capacityplanning

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplanning", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/capacityplanning")
	workforcemanagement_businessunits_capacityplanningCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplanning"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_capacityplanningCmd)
}

func Cmdworkforcemanagement_businessunits_capacityplanning() *cobra.Command {
	return workforcemanagement_businessunits_capacityplanningCmd
}
