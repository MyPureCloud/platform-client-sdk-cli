package workforcemanagement_agents_me_adherence

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("workforcemanagement_agents_me_adherence", "SWAGGER_OVERRIDE_/api/v2/workforcemanagement/agents/me/adherence")
	workforcemanagement_agents_me_adherenceCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("workforcemanagement_agents_me_adherence"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(workforcemanagement_agents_me_adherenceCmd)
}

func Cmdworkforcemanagement_agents_me_adherence() *cobra.Command {
	return workforcemanagement_agents_me_adherenceCmd
}
