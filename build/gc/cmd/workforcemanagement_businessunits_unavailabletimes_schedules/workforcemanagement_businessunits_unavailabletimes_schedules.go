package workforcemanagement_businessunits_unavailabletimes_schedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_unavailabletimes_schedules", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/unavailabletimes/schedules")
	workforcemanagement_businessunits_unavailabletimes_schedulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_unavailabletimes_schedules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_unavailabletimes_schedulesCmd)
}

func Cmdworkforcemanagement_businessunits_unavailabletimes_schedules() *cobra.Command {
	return workforcemanagement_businessunits_unavailabletimes_schedulesCmd
}
