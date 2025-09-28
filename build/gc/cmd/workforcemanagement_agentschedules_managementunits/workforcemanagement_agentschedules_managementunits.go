package workforcemanagement_agentschedules_managementunits

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agentschedules_managementunits", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agentschedules/managementunits")
	workforcemanagement_agentschedules_managementunitsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agentschedules_managementunits"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agentschedules_managementunitsCmd)
}

func Cmdworkforcemanagement_agentschedules_managementunits() *cobra.Command {
	return workforcemanagement_agentschedules_managementunitsCmd
}
