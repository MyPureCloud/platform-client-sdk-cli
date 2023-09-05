package workforcemanagement_managementunits_timeoffrequests_users

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_timeoffrequests_users", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/{timeOffRequestId}/users")
	workforcemanagement_managementunits_timeoffrequests_usersCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_timeoffrequests_users"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_timeoffrequests_usersCmd)
}

func Cmdworkforcemanagement_managementunits_timeoffrequests_users() *cobra.Command {
	return workforcemanagement_managementunits_timeoffrequests_usersCmd
}
