package workforcemanagement_agentschedules

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agentschedules", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agentschedules")
	workforcemanagement_agentschedulesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agentschedules"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agentschedulesCmd)
}

func Cmdworkforcemanagement_agentschedules() *cobra.Command {
	return workforcemanagement_agentschedulesCmd
}
