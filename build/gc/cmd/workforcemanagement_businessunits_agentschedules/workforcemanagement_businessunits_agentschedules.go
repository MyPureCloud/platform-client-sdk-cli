package workforcemanagement_businessunits_agentschedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_businessunits_agentschedules", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/businessunits/{businessUnitId}/agentschedules")
	workforcemanagement_businessunits_agentschedulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_businessunits_agentschedules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_businessunits_agentschedulesCmd)
}

func Cmdworkforcemanagement_businessunits_agentschedules() *cobra.Command {
	return workforcemanagement_businessunits_agentschedulesCmd
}
