package workforcemanagement_businessunits_unavailabletimes_settings

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_unavailabletimes_settings", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/unavailabletimes/settings")
	workforcemanagement_businessunits_unavailabletimes_settingsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_unavailabletimes_settings"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_unavailabletimes_settingsCmd)
}

func Cmdworkforcemanagement_businessunits_unavailabletimes_settings() *cobra.Command {
	return workforcemanagement_businessunits_unavailabletimes_settingsCmd
}
