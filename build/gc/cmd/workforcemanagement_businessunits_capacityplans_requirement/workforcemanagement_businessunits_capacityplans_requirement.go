package workforcemanagement_businessunits_capacityplans_requirement

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplans_requirement", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/capacityplans/{capacityPlanId}/requirement")
	workforcemanagement_businessunits_capacityplans_requirementCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_capacityplans_requirement"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_capacityplans_requirementCmd)
}

func Cmdworkforcemanagement_businessunits_capacityplans_requirement() *cobra.Command {
	return workforcemanagement_businessunits_capacityplans_requirementCmd
}
