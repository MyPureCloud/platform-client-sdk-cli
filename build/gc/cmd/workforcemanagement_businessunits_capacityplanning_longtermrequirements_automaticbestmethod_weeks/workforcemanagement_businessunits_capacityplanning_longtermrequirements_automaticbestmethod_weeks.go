package workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod_weeks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod_weeks", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/capacityplanning/longtermrequirements/automaticbestmethod/weeks")
	workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod_weeksCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod_weeks"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod_weeksCmd)
}

func Cmdworkforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod_weeks() *cobra.Command {
	return workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod_weeksCmd
}
