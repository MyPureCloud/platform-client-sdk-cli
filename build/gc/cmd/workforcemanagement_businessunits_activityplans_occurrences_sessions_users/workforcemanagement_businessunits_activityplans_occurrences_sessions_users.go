package workforcemanagement_businessunits_activityplans_occurrences_sessions_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_activityplans_occurrences_sessions_users", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}/occurrences/{occurrenceId}/sessions/{sessionId}/users")
	workforcemanagement_businessunits_activityplans_occurrences_sessions_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_activityplans_occurrences_sessions_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_activityplans_occurrences_sessions_usersCmd)
}

func Cmdworkforcemanagement_businessunits_activityplans_occurrences_sessions_users() *cobra.Command {
	return workforcemanagement_businessunits_activityplans_occurrences_sessions_usersCmd
}
