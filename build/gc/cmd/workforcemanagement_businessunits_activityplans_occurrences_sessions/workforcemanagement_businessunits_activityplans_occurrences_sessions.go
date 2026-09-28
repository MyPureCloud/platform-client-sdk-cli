package workforcemanagement_businessunits_activityplans_occurrences_sessions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_activityplans_occurrences_sessions", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}/occurrences/{occurrenceId}/sessions")
	workforcemanagement_businessunits_activityplans_occurrences_sessionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_activityplans_occurrences_sessions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_activityplans_occurrences_sessionsCmd)
}

func Cmdworkforcemanagement_businessunits_activityplans_occurrences_sessions() *cobra.Command {
	return workforcemanagement_businessunits_activityplans_occurrences_sessionsCmd
}
