package workforcemanagement_businessunits_capacityplanning_longtermrequirements

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplanning_longtermrequirements", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/capacityplanning/longtermrequirements")
	workforcemanagement_businessunits_capacityplanning_longtermrequirementsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplanning_longtermrequirements"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_capacityplanning_longtermrequirementsCmd)
}

func Cmdworkforcemanagement_businessunits_capacityplanning_longtermrequirements() *cobra.Command {
	return workforcemanagement_businessunits_capacityplanning_longtermrequirementsCmd
}
