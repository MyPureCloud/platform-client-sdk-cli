package workforcemanagement_managementunits_weeks_shifttrades_state

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_weeks_shifttrades_state", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state")
	workforcemanagement_managementunits_weeks_shifttrades_stateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_weeks_shifttrades_state"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_weeks_shifttrades_stateCmd)
}

func Cmdworkforcemanagement_managementunits_weeks_shifttrades_state() *cobra.Command {
	return workforcemanagement_managementunits_weeks_shifttrades_stateCmd
}
