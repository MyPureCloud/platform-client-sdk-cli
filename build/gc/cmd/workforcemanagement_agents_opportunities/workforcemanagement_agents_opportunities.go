package workforcemanagement_agents_opportunities

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_opportunities", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/opportunities")
	workforcemanagement_agents_opportunitiesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_opportunities"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_opportunitiesCmd)
}

func Cmdworkforcemanagement_agents_opportunities() *cobra.Command {
	return workforcemanagement_agents_opportunitiesCmd
}
