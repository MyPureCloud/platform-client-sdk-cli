package workforcemanagement_agents

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/me")
	workforcemanagement_agentsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agentsCmd)
}

func Cmdworkforcemanagement_agents() *cobra.Command {
	return workforcemanagement_agentsCmd
}
