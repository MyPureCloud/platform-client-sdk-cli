package workforcemanagement_businessunits_alternativeshifts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_alternativeshifts", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts")
	workforcemanagement_businessunits_alternativeshiftsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_alternativeshifts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_alternativeshiftsCmd)
}

func Cmdworkforcemanagement_businessunits_alternativeshifts() *cobra.Command {
	return workforcemanagement_businessunits_alternativeshiftsCmd
}
