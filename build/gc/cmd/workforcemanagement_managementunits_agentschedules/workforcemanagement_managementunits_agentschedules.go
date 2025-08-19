package workforcemanagement_managementunits_agentschedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_agentschedules", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/agentschedules")
	workforcemanagement_managementunits_agentschedulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_agentschedules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_agentschedulesCmd)
}

func Cmdworkforcemanagement_managementunits_agentschedules() *cobra.Command {
	return workforcemanagement_managementunits_agentschedulesCmd
}
