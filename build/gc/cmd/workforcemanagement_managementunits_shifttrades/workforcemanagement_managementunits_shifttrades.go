package workforcemanagement_managementunits_shifttrades

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_shifttrades", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades")
	workforcemanagement_managementunits_shifttradesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_shifttrades"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_shifttradesCmd)
}

func Cmdworkforcemanagement_managementunits_shifttrades() *cobra.Command {
	return workforcemanagement_managementunits_shifttradesCmd
}
