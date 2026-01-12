package workforcemanagement_agents_unavailabletimes

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_unavailabletimes", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/{agentId}/unavailabletimes")
	workforcemanagement_agents_unavailabletimesCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_unavailabletimes"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_unavailabletimesCmd)
}

func Cmdworkforcemanagement_agents_unavailabletimes() *cobra.Command {
	return workforcemanagement_agents_unavailabletimesCmd
}
