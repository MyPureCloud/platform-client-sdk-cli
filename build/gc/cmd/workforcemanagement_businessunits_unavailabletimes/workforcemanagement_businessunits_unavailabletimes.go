package workforcemanagement_businessunits_unavailabletimes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_unavailabletimes", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/unavailabletimes")
	workforcemanagement_businessunits_unavailabletimesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_unavailabletimes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_unavailabletimesCmd)
}

func Cmdworkforcemanagement_businessunits_unavailabletimes() *cobra.Command {
	return workforcemanagement_businessunits_unavailabletimesCmd
}
