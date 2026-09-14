package workforcemanagement_businessunits_schedulingpreferences

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_schedulingpreferences", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/schedulingpreferences")
	workforcemanagement_businessunits_schedulingpreferencesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_schedulingpreferences"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_schedulingpreferencesCmd)
}

func Cmdworkforcemanagement_businessunits_schedulingpreferences() *cobra.Command {
	return workforcemanagement_businessunits_schedulingpreferencesCmd
}
