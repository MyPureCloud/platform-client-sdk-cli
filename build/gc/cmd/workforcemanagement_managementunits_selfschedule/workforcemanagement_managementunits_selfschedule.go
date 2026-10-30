package workforcemanagement_managementunits_selfschedule

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_selfschedule", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/selfschedule")
	workforcemanagement_managementunits_selfscheduleCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_selfschedule"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_selfscheduleCmd)
}

func Cmdworkforcemanagement_managementunits_selfschedule() *cobra.Command {
	return workforcemanagement_managementunits_selfscheduleCmd
}
