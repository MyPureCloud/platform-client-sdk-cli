package workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/capacityplanning/longtermrequirements/automaticbestmethod")
	workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethodCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethodCmd)
}

func Cmdworkforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethod() *cobra.Command {
	return workforcemanagement_businessunits_capacityplanning_longtermrequirements_automaticbestmethodCmd
}
