package workforcemanagement_businessunits_mainforecast_continuousforecast

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_mainforecast_continuousforecast", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/mainforecast/continuousforecast")
	workforcemanagement_businessunits_mainforecast_continuousforecastCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_mainforecast_continuousforecast"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_mainforecast_continuousforecastCmd)
}

func Cmdworkforcemanagement_businessunits_mainforecast_continuousforecast() *cobra.Command {
	return workforcemanagement_businessunits_mainforecast_continuousforecastCmd
}
