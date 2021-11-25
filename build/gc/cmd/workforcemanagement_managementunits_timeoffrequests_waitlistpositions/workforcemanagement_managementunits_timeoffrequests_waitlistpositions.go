package workforcemanagement_managementunits_timeoffrequests_waitlistpositions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_timeoffrequests_waitlistpositions", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/waitlistpositions")
	workforcemanagement_managementunits_timeoffrequests_waitlistpositionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_timeoffrequests_waitlistpositions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_timeoffrequests_waitlistpositionsCmd)
}

func Cmdworkforcemanagement_managementunits_timeoffrequests_waitlistpositions() *cobra.Command {
	return workforcemanagement_managementunits_timeoffrequests_waitlistpositionsCmd
}
