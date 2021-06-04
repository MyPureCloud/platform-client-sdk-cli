package workforcemanagement_managementunits_schedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_schedules", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules")
	workforcemanagement_managementunits_schedulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_schedules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_schedulesCmd)
}

func Cmdworkforcemanagement_managementunits_schedules() *cobra.Command {
	return workforcemanagement_managementunits_schedulesCmd
}
