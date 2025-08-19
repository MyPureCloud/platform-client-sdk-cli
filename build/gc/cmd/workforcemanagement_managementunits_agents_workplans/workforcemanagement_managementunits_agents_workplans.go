package workforcemanagement_managementunits_agents_workplans

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_managementunits_agents_workplans", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/workplans")
	workforcemanagement_managementunits_agents_workplansCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_managementunits_agents_workplans"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_managementunits_agents_workplansCmd)
}

func Cmdworkforcemanagement_managementunits_agents_workplans() *cobra.Command {
	return workforcemanagement_managementunits_agents_workplansCmd
}
