package workforcemanagement_businessunits_mainforecast

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_mainforecast", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/mainforecast")
	workforcemanagement_businessunits_mainforecastCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_mainforecast"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_mainforecastCmd)
}

func Cmdworkforcemanagement_businessunits_mainforecast() *cobra.Command {
	return workforcemanagement_businessunits_mainforecastCmd
}
