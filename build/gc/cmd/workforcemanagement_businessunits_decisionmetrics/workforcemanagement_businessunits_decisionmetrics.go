package workforcemanagement_businessunits_decisionmetrics

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_decisionmetrics", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/decisionmetrics")
	workforcemanagement_businessunits_decisionmetricsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_decisionmetrics"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_decisionmetricsCmd)
}

func Cmdworkforcemanagement_businessunits_decisionmetrics() *cobra.Command {
	return workforcemanagement_businessunits_decisionmetricsCmd
}
