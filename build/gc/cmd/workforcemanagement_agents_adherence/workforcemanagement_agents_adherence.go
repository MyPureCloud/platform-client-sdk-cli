package workforcemanagement_agents_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_adherence", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/{agentId}/adherence")
	workforcemanagement_agents_adherenceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_adherence"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_adherenceCmd)
}

func Cmdworkforcemanagement_agents_adherence() *cobra.Command {
	return workforcemanagement_agents_adherenceCmd
}
