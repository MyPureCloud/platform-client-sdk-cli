package workforcemanagement_businessunits_minimumstaffing

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_minimumstaffing", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/minimumstaffing")
	workforcemanagement_businessunits_minimumstaffingCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_minimumstaffing"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_minimumstaffingCmd)
}

func Cmdworkforcemanagement_businessunits_minimumstaffing() *cobra.Command {
	return workforcemanagement_businessunits_minimumstaffingCmd
}
