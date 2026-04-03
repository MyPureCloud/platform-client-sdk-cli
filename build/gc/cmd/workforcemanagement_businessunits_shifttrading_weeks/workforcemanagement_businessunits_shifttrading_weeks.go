package workforcemanagement_businessunits_shifttrading_weeks

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_weeks", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/shifttrading/weeks")
	workforcemanagement_businessunits_shifttrading_weeksCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_shifttrading_weeks"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_shifttrading_weeksCmd)
}

func Cmdworkforcemanagement_businessunits_shifttrading_weeks() *cobra.Command {
	return workforcemanagement_businessunits_shifttrading_weeksCmd
}
