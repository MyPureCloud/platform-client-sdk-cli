package workforcemanagement_managementunits_timeoffrequests_integrationstatus

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_timeoffrequests_integrationstatus", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/integrationstatus")
	workforcemanagement_managementunits_timeoffrequests_integrationstatusCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_timeoffrequests_integrationstatus"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_timeoffrequests_integrationstatusCmd)
}

func Cmdworkforcemanagement_managementunits_timeoffrequests_integrationstatus() *cobra.Command {
	return workforcemanagement_managementunits_timeoffrequests_integrationstatusCmd
}
