package workforcemanagement_managementunits_unavailabletimes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_unavailabletimes", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/unavailabletimes")
	workforcemanagement_managementunits_unavailabletimesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_unavailabletimes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_unavailabletimesCmd)
}

func Cmdworkforcemanagement_managementunits_unavailabletimes() *cobra.Command {
	return workforcemanagement_managementunits_unavailabletimesCmd
}
