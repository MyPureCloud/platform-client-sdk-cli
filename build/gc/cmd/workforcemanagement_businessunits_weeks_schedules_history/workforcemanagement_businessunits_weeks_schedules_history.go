package workforcemanagement_businessunits_weeks_schedules_history

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_schedules_history", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history")
	workforcemanagement_businessunits_weeks_schedules_historyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_schedules_history"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_weeks_schedules_historyCmd)
}

func Cmdworkforcemanagement_businessunits_weeks_schedules_history() *cobra.Command {
	return workforcemanagement_businessunits_weeks_schedules_historyCmd
}
