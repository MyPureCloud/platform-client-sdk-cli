package workforcemanagement_businessunits_weeks_schedules_agentschedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_schedules_agentschedules", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules")
	workforcemanagement_businessunits_weeks_schedules_agentschedulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_weeks_schedules_agentschedules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_weeks_schedules_agentschedulesCmd)
}

func Cmdworkforcemanagement_businessunits_weeks_schedules_agentschedules() *cobra.Command {
	return workforcemanagement_businessunits_weeks_schedules_agentschedulesCmd
}
