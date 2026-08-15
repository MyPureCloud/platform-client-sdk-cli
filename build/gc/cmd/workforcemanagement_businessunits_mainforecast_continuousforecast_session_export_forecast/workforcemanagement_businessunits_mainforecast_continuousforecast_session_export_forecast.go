package workforcemanagement_businessunits_mainforecast_continuousforecast_session_export_forecast

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_mainforecast_continuousforecast_session_export_forecast", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/mainforecast/continuousforecast/session/export/forecast")
	workforcemanagement_businessunits_mainforecast_continuousforecast_session_export_forecastCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_mainforecast_continuousforecast_session_export_forecast"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_mainforecast_continuousforecast_session_export_forecastCmd)
}

func Cmdworkforcemanagement_businessunits_mainforecast_continuousforecast_session_export_forecast() *cobra.Command {
	return workforcemanagement_businessunits_mainforecast_continuousforecast_session_export_forecastCmd
}
