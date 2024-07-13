package workforcemanagement_businessunits_activityplans_runs

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_activityplans_runs", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}/runs")
	workforcemanagement_businessunits_activityplans_runsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_activityplans_runs"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_activityplans_runsCmd)
}

func Cmdworkforcemanagement_businessunits_activityplans_runs() *cobra.Command {
	return workforcemanagement_businessunits_activityplans_runsCmd
}
