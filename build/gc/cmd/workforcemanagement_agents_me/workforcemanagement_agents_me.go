package workforcemanagement_agents_me

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_me", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/me")
	workforcemanagement_agents_meCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_me"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_meCmd)
}

func Cmdworkforcemanagement_agents_me() *cobra.Command {
	return workforcemanagement_agents_meCmd
}
